import * as vscode from 'vscode';
import * as cp from 'child_process';
import * as path from 'path';
import * as fs from 'fs';

function resolveBinaryName(): string {
    const platform = process.platform;
    const arch = process.arch;

    if (platform === 'win32') {
        return 'paste-go.exe';
    }

    if (platform === 'darwin') {
        return arch === 'arm64' ? 'paste-go-darwin-arm64' : 'paste-go-darwin-amd64';
    }

    if (platform === 'linux') {
        return arch === 'arm64' ? 'paste-go-linux-arm64' : 'paste-go-linux-amd64';
    }

    return 'paste-go';
}

function findBundledBinary(context: vscode.ExtensionContext): string | null {
    const binName = resolveBinaryName();

    const candidates = [
        path.join(context.extensionPath, 'bin', binName)
    ];

    const installed = vscode.extensions.getExtension('cointem.paste-go');
    if (installed && installed.extensionPath && installed.extensionPath !== context.extensionPath) {
        candidates.push(path.join(installed.extensionPath, 'bin', binName));
    }

    // also try legacy locations just in case
    candidates.push(path.join(context.extensionPath, '..', 'extension', 'bin', binName));
    candidates.push(path.join(context.extensionPath, '..', '..', 'extension', 'bin', binName));

    for (const c of candidates) {
        try {
            if (fs.existsSync(c)) return c;
        } catch (e) {
            // ignore permission errors and continue
        }
    }
    return null;
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

           const config = vscode.workspace.getConfiguration('pasteGo');
        const aiKey = config.get<string>('aiApiKey');
        const aiApiFormat = config.get<string>('aiApiFormat') || "gemini";
        const aiModel = config.get<string>('aiModel') || "";
        const aiBaseUrl = config.get<string>('aiBaseUrl') || "";

        // Resolve binary strictly from bundled/installed extension
        const binPath = findBundledBinary(context);
        outputChannel.appendLine(`Core Path: ${binPath ?? '(not found)'}`);

        if (!binPath) {
            outputChannel.appendLine('Error: paste-go binary not found in extension bin folder.');
            outputChannel.appendLine('Checked in extensionPath and installed extension locations.');
            outputChannel.show(true);
            vscode.window.showErrorMessage('Paste Go core binary not found. Reinstall the extension or set a valid path in settings.');
            return;
        }

        // Prepare args
        const languageId = editor.document.languageId;
        const args = ['-lang', languageId];
        if (aiKey) args.push('-key', aiKey);
        if (aiApiFormat) args.push('-format', aiApiFormat);
        if (aiModel) args.push('-model', aiModel);
        if (aiBaseUrl) args.push('-baseurl', aiBaseUrl);

        outputChannel.appendLine(`Spawning with args: ${args.join(' ')}`);



        let output = '';
        let errorOutput = '';

        const proc = cp.spawn(binPath, args);

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
