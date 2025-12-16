document.addEventListener('DOMContentLoaded', function() {
    checkAuth();
});

function checkAuth() {
    const user = localStorage.getItem('user');
    const currentPage = window.location.pathname.split('/').pop();
    
    if (currentPage === 'login.html' && user) {
        window.location.href = 'index.html';
    }
    
    if (currentPage !== 'login.html' && !user) {
        window.location.href = 'login.html';
    }
}

function login() {
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    
    if (email && password) {
        const user = {
            name: "Иванов Иван Иванович",
            email: email,
            group: "ФТ-101"
        };
        
        localStorage.setItem('user', JSON.stringify(user));
        window.location.href = 'index.html';
    } else {
        alert("Введите email и пароль");
    }
}

function logout() {
    localStorage.removeItem('user');
    window.location.href = 'login.html';
}