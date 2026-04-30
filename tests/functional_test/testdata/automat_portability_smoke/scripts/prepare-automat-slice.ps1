param(
    [string]$DependencyContract = "portable-dependencies.json",
    [string]$WorkflowGuide = "docs/portable-workflow.md"
)

if (-not (Test-Path -LiteralPath $DependencyContract)) {
    throw "missing dependency contract: $DependencyContract"
}

if (-not (Test-Path -LiteralPath $WorkflowGuide)) {
    throw "missing workflow guide: $WorkflowGuide"
}

$guide = Get-Content -Raw -LiteralPath $WorkflowGuide
$contract = Get-Content -Raw -LiteralPath $DependencyContract | ConvertFrom-Json
$requiredTools = @($contract.requiredTools | ForEach-Object { $_.name })

if (-not $guide.Contains("Portable Workflow Slice")) {
    throw "workflow guide missing portability heading"
}

foreach ($toolName in $requiredTools) {
    if (-not $guide.Contains($toolName)) {
        throw "workflow guide missing declared tool: $toolName"
    }
}

$guideHeading = (($guide -split "`r?`n") | Where-Object { $_.Trim() -ne "" } | Select-Object -First 1).TrimStart("# ").Trim()

Write-Output ("dispatch-ready:" + $guideHeading + ":" + ($requiredTools -join ","))
