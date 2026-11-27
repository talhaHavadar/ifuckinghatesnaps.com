// The theme picker. The choice is system (follow the OS), light or dark; the
// head script has already applied a saved light or dark choice before the
// page painted. Without this script the picker stays hidden and the page
// follows the system through CSS alone.
(() => {
  const button = document.querySelector(".theme");
  if (!button) return;
  const root = document.documentElement;
  const order = ["system", "light", "dark"];
  const names = { system: "follows your system", light: "light", dark: "dark" };

  const read = () => {
    try {
      const saved = localStorage.getItem("theme");
      return saved === "light" || saved === "dark" ? saved : "system";
    } catch {
      return "system";
    }
  };

  const apply = (choice) => {
    if (choice === "system") delete root.dataset.theme;
    else root.dataset.theme = choice;
    button.dataset.choice = choice;
    button.setAttribute("aria-label", `Colour theme: ${names[choice]}. Change it.`);
    button.title = `Theme: ${names[choice]}`;
  };

  apply(read());
  button.hidden = false;
  button.addEventListener("click", () => {
    const next = order[(order.indexOf(button.dataset.choice) + 1) % order.length];
    try {
      if (next === "system") localStorage.removeItem("theme");
      else localStorage.setItem("theme", next);
    } catch {
      // Without storage the choice lasts until the page is left.
    }
    apply(next);
  });
})();
