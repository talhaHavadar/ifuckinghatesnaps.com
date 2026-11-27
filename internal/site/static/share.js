// The share button on a quote page. Without this script the quote's link is
// shown as selectable text and the button stays hidden. On a phone it opens
// the system share sheet; elsewhere it copies the link.
(() => {
  const button = document.querySelector(".share");
  if (!button) return;
  const status = document.getElementById("share-status");
  const url = button.dataset.url;
  const canShare = typeof navigator.share === "function" && matchMedia("(pointer: coarse)").matches;
  const label = canShare ? "Share" : "Copy link";
  button.textContent = label;
  button.hidden = false;

  let reset = 0;
  const say = (text) => {
    button.textContent = text;
    if (status) status.textContent = text;
    clearTimeout(reset);
    reset = setTimeout(() => {
      button.textContent = label;
    }, 1800);
  };

  // Without clipboard access, select the visible link so it can be copied by hand.
  const selectLink = () => {
    const link = document.querySelector(".permalink");
    if (!link) return;
    const range = document.createRange();
    range.selectNodeContents(link);
    const selection = getSelection();
    selection.removeAllRanges();
    selection.addRange(range);
  };

  button.addEventListener("click", async () => {
    if (canShare) {
      try {
        await navigator.share({ title: button.dataset.title, url });
      } catch {
        // Closing the share sheet is not an error worth reporting.
      }
      return;
    }
    try {
      await navigator.clipboard.writeText(url);
      say("Copied");
    } catch {
      selectLink();
      say("Link selected");
    }
  });
})();
