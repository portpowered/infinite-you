param(
    [string]$DependencyContract = "portable-dependencies.json"
)

$contract = Get-Content -Raw -LiteralPath $DependencyContract | ConvertFrom-Json
$names = @($contract.requiredTools | ForEach-Object { $_.name })

Write-Output ("required-tools:" + ($names -join ","))
