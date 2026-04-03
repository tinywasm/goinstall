$version = if ($args[0]) { $args[0] } else { "1.25.2" }
$url = "https://go.dev/dl/go$version.windows-amd64.msi"
$output = "$env:TEMP\go.msi"

Write-Host "Downloading Go $version..."
Invoke-WebRequest -Uri $url -OutFile $output

Write-Host "Installing Go..."
Start-Process msiexec.exe -Wait -ArgumentList "/i `"$output`" /quiet /norestart"

Write-Host "Verifying installation..."
& "C:\Program Files\Go\bin\go.exe" version

Remove-Item $output
Write-Host "Go $version installed successfully."
