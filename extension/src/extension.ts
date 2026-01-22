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

function ensureExecutable(filePath: string, outputChannel: vscode.OutputChannel) {
    if (process.platform === 'win32') {
        return;
    }

    try {
        const stat = fs.statSync(filePath);
        const isExecutable = (stat.mode & 0o111) !== 0;
        if (!isExecutable) {
            fs.chmodSync(filePath, 0o755);
            outputChannel.appendLine(`chmod +x applied: ${filePath}`);
        }
    } catch (error) {
        outputChannel.appendLine(`chmod check failed: ${error}`);
    }
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
        const aiKey = config.get<string>('aiApiKey');
        const aiApiFormat = config.get<string>('aiApiFormat') || "gemini";
        const aiModel = config.get<string>('aiModel') || "";
        const aiBaseUrl = config.get<string>('aiBaseUrl') || "";

        const extPath = context.extensionPath;
        const binPath = resolveBundledBinary(extPath);

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
        outputChannel.appendLine(`Core Path: ${binPath}`);
        if (!fs.existsSync(binPath)) {
            outputChannel.show(true);
            vscode.window.showErrorMessage('Paste Go binary not found. Please reinstall the extension.');
            return;
        }

        ensureExecutable(binPath, outputChannel);
        const proc: cp.ChildProcess = cp.spawn(binPath, args);
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
                const trimmedError = errorOutput.trim();
                const lastErrorLine = trimmedError.split('\n').slice(-1)[0] || 'Unknown error';
                const hintParts: string[] = [];

                if (/permission denied/i.test(trimmedError)) {
                    hintParts.push('Binary lacks execute permission. Try reinstalling or run chmod +x on the bundled binary.');
                }
                if (/401|unauthorized/i.test(trimmedError)) {
                    hintParts.push('Unauthorized. Check pasteGo.aiApiKey.');
                }
                if (/403|forbidden/i.test(trimmedError)) {
                    hintParts.push('Forbidden. Check API key permissions or IP restrictions.');
                }
                if (/429|rate limit/i.test(trimmedError)) {
                    hintParts.push('Rate limited. Wait and retry, or switch provider.');
                }
                if (/timeout|timed out|context deadline exceeded/i.test(trimmedError)) {
                    hintParts.push('Network timeout. Check proxy or pasteGo.aiBaseUrl.');
                }
                if (/no such file or directory/i.test(trimmedError)) {
                    hintParts.push('Binary not found. Reinstall the extension.');
                }

                const hintText = hintParts.length > 0 ? `\nHint: ${hintParts.join(' ')}` : '';
                outputChannel.show(true);
                vscode.window.showErrorMessage(`Paste Go failed (code ${code}). ${lastErrorLine}${hintText}`);
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