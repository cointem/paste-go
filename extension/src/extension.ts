import * as vscode from 'vscode';
import * as cp from 'child_process';
import * as path from 'path';
import * as fs from 'fs';

function resolveBundledBinary(extPath: string): string {
    const platform = process.platform;
    const arch = process.arch;

    if (platform === 'win32') {
        return path.join(extPath, 'bin', 'paste-go.exe');
    }

    if (platform === 'darwin') {
        if (arch === 'arm64') {
            return path.join(extPath, 'bin', 'paste-go-darwin-arm64');
        }
        return path.join(extPath, 'bin', 'paste-go-darwin-amd64');
    }

    if (platform === 'linux') {
        if (arch === 'arm64') {
            return path.join(extPath, 'bin', 'paste-go-linux-arm64');
        }
        return path.join(extPath, 'bin', 'paste-go-linux-amd64');
    }

    return path.join(extPath, 'bin', 'paste-go');
}

export function activate(context: vscode.ExtensionContext) {
    console.log('Paste Go is now active!');
    const outputChannel = vscode.window.createOutputChannel("Paste Go");

    let disposable = vscode.commands.registerCommand('paste-go.smartPaste', async () => {
        outputChannel.clear();
        outputChannel.appendLine("Smart Paste triggered");

        const editor = vscode.window.activeTextEditor;
        if (!editor) {
            return;
        }

        // 1. Get Clipboard Content
        const clipboardText = await vscode.env.clipboard.readText();
        outputChannel.appendLine(`Clipboard length: ${clipboardText.length}`);
        if (!clipboardText) {
            vscode.window.showInformationMessage('Clipboard is empty');
            return;
        }

        // 2. Resolve Binary Path
        const config = vscode.workspace.getConfiguration('pasteGo');
        let binPath = config.get<string>('corePath');
        const aiKey = config.get<string>('aiApiKey');
        const aiApiFormat = config.get<string>('aiApiFormat') || "gemini";
        const aiModel = config.get<string>('aiModel') || "";
        const aiBaseUrl = config.get<string>('aiBaseUrl') || "";

        if (!binPath) {
            const extPath = context.extensionPath;
            binPath = resolveBundledBinary(extPath);
        }

        // 3. Prepare Arguments
        // We detect the language of the current file to pass as target
        const languageId = editor.document.languageId; 
        const args = ['-lang', languageId];
        if (aiKey) {
            args.push('-key', aiKey);
        }
        if (aiApiFormat) {
            args.push('-format', aiApiFormat);
        }
        if (aiModel) {
            args.push('-model', aiModel);
        }
        if (aiBaseUrl) {
            args.push('-baseurl', aiBaseUrl);
        }

        // 4. Execute Core
        // Quick dev hack: if binPath doesn't exist, try running with 'go run' from workspace if available
        // But for a robust extension, we should valid binary. 
        // For THIS DEMO session, to make it work immediately without compiling binary:
        let proc: cp.ChildProcess;
        
        // CHECK IF DEV MODE: If we are in the dev workspace, we can use `go run`
        outputChannel.appendLine(`Core Path: ${binPath}`);
        
        if (binPath.includes('paste-go') && !fs.existsSync(binPath)) {
             // Fallback to go run for development convenience
             outputChannel.appendLine("Binary not found, falling back to 'go run'...");
             const coreDir = path.join(context.extensionPath, '..', 'core');
             proc = cp.spawn('go', ['run', './cmd/paste-go/main.go', ...args], {
                 cwd: coreDir,
                 env: process.env // Inherit env for GOPATH etc
             });
        } else {
             proc = cp.spawn(binPath, args);
        }
        
        outputChannel.appendLine(`Spawning with args: ${args.join(' ')}`);

        let output = '';
        let errorOutput = '';

        if (proc.stdin) {
            try {
                proc.stdin.write(clipboardText);
                proc.stdin.end();
            } catch (e) {
                outputChannel.appendLine(`Error writing to stdin: ${e}`);
            }
        }

        proc.stdout?.on('data', (data) => {
            output += data.toString();
        });

        proc.stderr?.on('data', (data) => {
            errorOutput += data.toString();
            outputChannel.appendLine(`STDERR: ${data.toString()}`);
        });

        proc.on('close', (code) => {
            outputChannel.appendLine(`Process exited with code: ${code}`);
            if (code !== 0) {
                outputChannel.show(true); // Show output channel on error
                vscode.window.showErrorMessage(`Paste Forge failed. Check Output channel for details.`);
                return;
            }

            if (!output) {
                outputChannel.appendLine("Warning: Output is empty!");
                vscode.window.showWarningMessage("Paste Forge returned no content.");
            }

            // 5. Insert Result
            editor.edit(editBuilder => {
                editor.selections.forEach(selection => {
                    if (selection.isEmpty) {
                        editBuilder.insert(selection.active, output);
                    } else {
                        editBuilder.replace(selection, output);
                    }
                });
            });
        });
    });

    context.subscriptions.push(disposable);
}

export function deactivate() {}
