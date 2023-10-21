document
  .getElementById("btn-github-com-unset-width")
  .addEventListener("click", async () => {
    let [tab] = await chrome.tabs.query({ active: true, currentWindow: true });
    chrome.scripting.executeScript({
      target: { tabId: tab.id },
      function: onRun,
    });
  });

function onRun() {
  const els = document.getElementsByClassName("Layout-sidebar");
  if (els.length < 1) {
    return;
  }
  el = els[0];
  el.style.width = "unset";
}
