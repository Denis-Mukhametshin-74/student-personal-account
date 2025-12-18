export async function login() {
    const email = document.getElementById('email').value;
    const password = document.getElementById('password').value;
    
    if (!email || !password) {
        alert("Введите email и пароль");
        return;
    }
    
    try {
        const response = await fetch('/api/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                email: email,
                password: password
            })
        });
        
        if (response.ok) {
            const data = await response.json();
            localStorage.setItem('token', data.token);
            localStorage.setItem('user', JSON.stringify(data.user));
            window.location.href = 'index.html';
        } else {
            alert("Ошибка авторизации. Проверьте email и пароль.");
        }
    } catch (error) {
        console.error('Login error:', error);
        alert("Ошибка соединения с сервером");
    }
}

export async function register() {
    const name = document.getElementById('regName').value;
    const email = document.getElementById('regEmail').value;
    const group = document.getElementById('regGroup').value;
    const password = document.getElementById('regPassword').value;
    const confirmPassword = document.getElementById('regConfirmPassword').value;
    
    if (!name || !email || !group || !password || !confirmPassword) {
        alert("Заполните все поля");
        return;
    }
    
    if (password !== confirmPassword) {
        alert("Пароли не совпадают");
        return;
    }
    
    try {
        const response = await fetch('/api/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: name,
                email: email,
                group: group,
                password: password
            })
        });
        
        if (response.ok) {
            alert("Регистрация успешна! Теперь войдите в систему.");
            window.location.href = 'login.html';
        } else {
            const error = await response.text();
            alert("Ошибка регистрации: " + error);
        }
    } catch (error) {
        console.error('Register error:', error);
        alert("Ошибка соединения с сервером");
    }
}

export function logout() {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    window.location.href = 'login.html';
}