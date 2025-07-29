Get-Content .env | ForEach-Object {
    if ($_ -match "^\s*#") { return }    # Skip comments
    if ($_ -match "^\s*$") { return }    # Skip empty lines

    if ($_ -notmatch "=") { return }     # Skip lines without '='

    $parts = $_ -split '=', 2
    if ($parts.Count -lt 2) { return }   # Ensure both name and value are present

    $name = $parts[0].Trim()
    $value = $parts[1].Trim().Trim("'`"")

    Set-Item -Path "env:$name" -Value $value
}

go run main.go