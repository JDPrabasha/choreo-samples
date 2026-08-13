const btn = document.getElementById('generate-btn');
const status = document.getElementById('status');

btn.addEventListener('click', async () => {
  btn.disabled = true;
  status.textContent = 'Generating report... this can take a while.';
  const start = Date.now();

  try {
    const res = await fetch('/api/generate-report');
    if (!res.ok) {
      throw new Error(`HTTP ${res.status}`);
    }
    const data = await res.json();
    status.textContent = data.message;
  } catch (err) {
    const elapsed = ((Date.now() - start) / 1000).toFixed(1);
    status.textContent = `Request failed after ${elapsed}s: ${err.message}`;
  } finally {
    btn.disabled = false;
  }
});
