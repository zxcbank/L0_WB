document.getElementById('searchForm').addEventListener('submit', async (e) => {
    e.preventDefault();
    const formData = new FormData(e.target);
    const response = await fetch('/api/order', {
        method: 'POST',
        body: formData
    });
    const result = await response.json();
});