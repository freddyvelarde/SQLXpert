export function emptySpaceValidation(str: string): boolean {
  for (const char of str) {
    if (char === " ") return false;
  }
  return true;
}
