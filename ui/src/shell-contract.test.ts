import { readFileSync } from "node:fs";
import path from "node:path";
import { fileURLToPath } from "node:url";

const sourceDir = path.dirname(fileURLToPath(import.meta.url));
const uiDir = path.resolve(sourceDir, "..");

function readShell(relativePath: string): string {
  return readFileSync(path.join(uiDir, relativePath), "utf8");
}

describe("dashboard shell contract", () => {
  it("publishes Infinite You metadata and iconography in the live shell", () => {
    const shell = readShell("index.html");

    expect(shell).toContain("<title>Infinite You Dashboard</title>");
    expect(shell).toContain(
      'content="Standalone live dashboard shell for Infinite You."',
    );
    expect(shell).toContain("Infinite%20You%20dashboard%20icon");
    expect(shell).toContain('<div id="root"></div>');
    expect(shell).not.toContain("Agent Factory Dashboard");
    expect(shell).not.toContain(
      "Standalone live dashboard shell for Agent Factory.",
    );
  });

  it("keeps the checked-in fallback shell aligned with the live shell branding", () => {
    const shell = readShell("fallback_dist/index.html");

    expect(shell).toContain("<title>Infinite You Dashboard</title>");
    expect(shell).toContain(
      'content="Standalone live dashboard shell for Infinite You."',
    );
    expect(shell).toContain("Infinite%20You%20dashboard%20icon");
    expect(shell).toContain('<div id="root"></div>');
    expect(shell).not.toContain("Agent Factory Dashboard");
  });
});
