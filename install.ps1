# One-shot install: download latest chatsafe, put it in ~\bin, add to user PATH.
$ErrorActionPreference = "Stop"

$BinDir = Join-Path $HOME "bin"
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$Base = "https://github.com/sshehrozali/chatsafe/releases/latest/download"
$Zip = Join-Path $BinDir "chatsafe.zip"
Write-Host "Downloading chatsafe-windows-amd64.zip ..."
Invoke-WebRequest -Uri "$Base/chatsafe-windows-amd64.zip" -OutFile $Zip
Expand-Archive -Path $Zip -DestinationPath $BinDir -Force
Remove-Item $Zip

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$BinDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$BinDir", "User")
    Write-Host "Added $BinDir to your user PATH."
} else {
    Write-Host "PATH already includes $BinDir (skipped)."
}

Write-Host ""
Write-Host "Done. Open a new PowerShell window, then run: chatsafe.exe -version"
