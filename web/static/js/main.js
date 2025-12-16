document.addEventListener('DOMContentLoaded', async function() {
    await checkAuth();
});

async function checkAuth() {
    const token = localStorage.getItem('token');
    const currentPage = window.location.pathname.split('/').pop();
    
    if ((currentPage === 'login.html' || currentPage === 'register.html') && token) {
        window.location.href = 'index.html';
        return;
    }
    
    if (currentPage !== 'login.html' && currentPage !== 'register.html' && !token) {
        window.location.href = 'login.html';
        return;
    }
}