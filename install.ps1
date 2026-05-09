# One-shot install: download latest llmsave, put it in ~\bin, add to user PATH.
$ErrorActionPreference = "Stop"

$BinDir = Join-Path $HOME "bin"
New-Item -ItemType Directory -Force -Path $BinDir | Out-Null

$Zip = Join-Path $BinDir "llmsave.zip"

Write-Host "Fetching latest release for windows/amd64..."
$headers = @{ "User-Agent" = "llmsave-install" }
if ($env:GITHUB_TOKEN) {
    $headers["Authorization"] = "Bearer $($env:GITHUB_TOKEN)"
}
$rel = Invoke-RestMethod -Uri "https://api.github.com/repos/sshehrozali/llmsave/releases/latest" -Headers $headers

$asset = $rel.assets | Where-Object { $_.name -eq "llmsave-windows-amd64.zip" } | Select-Object -First 1
if (-not $asset) {
    $asset = $rel.assets | Where-Object { $_.name -match '^llmsave_[0-9.]+_windows_amd64\.zip$' } | Select-Object -First 1
}
if (-not $asset) {
    throw "No Windows amd64 zip in latest release. See https://github.com/sshehrozali/llmsave/releases"
}

Write-Host "Downloading..."
Invoke-WebRequest -Uri $asset.browser_download_url -OutFile $Zip

Expand-Archive -Path $Zip -DestinationPath $BinDir -Force
Remove-Item $Zip

$exe = Join-Path $BinDir "llmsave.exe"
if (-not (Test-Path $exe)) {
    throw "Extracted archive did not contain llmsave.exe"
}

$UserPath = [Environment]::GetEnvironmentVariable("Path", "User")
if ($UserPath -notlike "*$BinDir*") {
    [Environment]::SetEnvironmentVariable("Path", "$UserPath;$BinDir", "User")
    Write-Host "Added $BinDir to your user PATH."
} else {
    Write-Host "PATH already includes $BinDir (skipped)."
}

Write-Host ""
Write-Host "Done. Open a new PowerShell window, then run: llmsave.exe -version"
Write-Host "Create a backup: llmsave backup"
