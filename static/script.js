const API_BASE_URL = window.location.origin;
const FORM_ENDPOINT = `${API_BASE_URL}/api/url`;

const urlForm = document.getElementById('urlForm');
const resultDiv = document.getElementById('result');
const errorDiv = document.getElementById('error');
const shortUrlInput = document.getElementById('shortUrl');
const originalUrlSpan = document.getElementById('originalUrl');
const copyBtn = document.getElementById('copyBtn');
const submitBtn = document.getElementById('submitBtn');

function showToast(message) {
    const toast = document.getElementById('toast');
    toast.textContent = message;
    toast.style.display = 'block';
    setTimeout(() => {
        toast.style.display = 'none';
    }, 2000);
}

async function copyToClipboard(text) {
    try {
        await navigator.clipboard.writeText(text);
        showToast('Copied!');
        return true;
    } catch (err) {
        showToast('Failed to copy');
        return false;
    }
}

function showError(message) {
    const errorMessageSpan = document.getElementById('errorMessage');
    errorMessageSpan.textContent = message;
    errorDiv.style.display = 'flex';
    resultDiv.style.display = 'none';
    setTimeout(() => {
        errorDiv.style.display = 'none';
    }, 5000);
}

function showResult(shortUrl, originalUrl) {
    const fullShortUrl = `${API_BASE_URL}/${shortUrl}`;
    shortUrlInput.value = fullShortUrl;
    originalUrlSpan.textContent = originalUrl;
    resultDiv.style.display = 'block';
    errorDiv.style.display = 'none';
}

async function shortenUrl(url, alias) {
    const requestBody = { url };
    if (alias) requestBody.alias = alias;
    const response = await fetch(FORM_ENDPOINT, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(requestBody),
    });
    if (!response.ok) {
        const error = await response.json();
        throw new Error(error.error || 'Failed to shorten URL');
    }
    return await response.json();
}

urlForm.addEventListener('submit', async (e) => {
    e.preventDefault();
    const urlInput = document.getElementById('url');
    const aliasInput = document.getElementById('alias');
    const url = urlInput.value.trim();
    const alias = aliasInput.value.trim();
    if (!url) {
        showError('Please enter a URL');
        urlInput.classList.add('error');
        return;
    }
    urlInput.classList.remove('error');
    aliasInput.classList.remove('error');
    submitBtn.disabled = true;
    try {
        const data = await shortenUrl(url, alias);
        showResult(data.alias, url);
        urlForm.reset();
        showToast('URL shortened!');
    } catch (error) {
        showError(error.message);
    } finally {
        submitBtn.disabled = false;
    }
});

copyBtn.addEventListener('click', () => {
    if (shortUrlInput.value) {
        copyToClipboard(shortUrlInput.value);
    }
});

shortUrlInput.addEventListener('click', () => {
    if (shortUrlInput.value) {
        copyToClipboard(shortUrlInput.value);
    }
});