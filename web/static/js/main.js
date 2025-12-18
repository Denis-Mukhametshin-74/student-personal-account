import { login, register, logout } from "./auth.js";

document.addEventListener('DOMContentLoaded', async function() {
    const loginBtn = document.querySelector('.login-button');
    const registerBtn = document.querySelector('.register-button');
    const logoutBtn = document.querySelector('.loguot-button');

    if (loginBtn) loginBtn.addEventListener('click', login);
    if (registerBtn) registerBtn.addEventListener('click', register);
    if (logoutBtn) logoutBtn.addEventListener('click', logout);

    await checkAuth();

    if (window.location.pathname.includes('index.html')) {
        await loadSchedule();
    }
    if (window.location.pathname.includes('profile.html')) {
        await loadProfile();
        await loadSubjects();
    }
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

async function loadSchedule() {
    try {
        const token = localStorage.getItem('token');
        const response = await fetch('/api/schedule', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        if (response.ok) {
            const scheduleData = await response.json();
            renderSchedule(scheduleData);
        } else {
            console.log('Ошибка получения данных о расписании');
        }
    } catch (error) {
        console.error('Failed to load schedule:', error);
    }
}

async function loadProfile() {
    try {
        const token = localStorage.getItem('token');
        const response = await fetch('/api/profile', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        if (response.ok) {
            const profile = await response.json();
            renderProfile(profile);
        } else {
            console.log('Ошибка получения данных о профиле');
        }
    } catch (error) {
        console.error('Failed to load profile:', error);
    }
}

async function loadSubjects() {
    try {
        const token = localStorage.getItem('token');
        const response = await fetch('/api/subjects', {
            headers: {
                'Authorization': `Bearer ${token}`
            }
        });
        
        if (response.ok) {
            const subjects = await response.json();
            renderSubjects(subjects);
        } else {
            console.log('Ошибка получения данных о предметах');
        }
    } catch (error) {
        console.error('Failed to load subjects:', error);
    }
}

function renderSchedule(scheduleData) {
    const title = document.getElementById('pageTitle');
    const container = document.getElementById('scheduleContainer');
    
    if (title && scheduleData.group) {
        title.textContent = `Расписание группы ${scheduleData.group}`;
    }
    
    if (container && scheduleData.days) {
        container.innerHTML = '';
        
        scheduleData.days.forEach(day => {
            const dayElement = document.createElement('div');
            dayElement.className = 'day';
            
            const dayTitle = document.createElement('h3');
            dayTitle.textContent = day.day;
            dayElement.appendChild(dayTitle);
            
            day.lessons.forEach(lesson => {
                const lessonElement = document.createElement('div');
                lessonElement.className = 'lesson';
                lessonElement.innerHTML = `
                    <span class="time">${lesson.time}</span>
                    <span class="subject">${lesson.subject}</span>
                    <span class="room">${lesson.room}</span>
                    <span class="teacher">${lesson.teacher}</span>
                `;
                dayElement.appendChild(lessonElement);
            });
            
            container.appendChild(dayElement);
        });
    }
}

function renderProfile(profile) {
    const container = document.getElementById('profileInfo');
    if (container) {
        container.innerHTML = `
            <p><strong>ФИО:</strong> ${profile.name}</p>
            <p><strong>Группа:</strong> ${profile.group}</p>
            <p><strong>Email:</strong> ${profile.email}</p>
        `;
    }
}

function renderSubjects(subjects) {
    const container = document.getElementById('subjectsContainer');
    if (container && subjects.length > 0) {
        container.innerHTML = subjects.map(subject => `
            <div class="subject-card">
                <h3>${subject.name}</h3>
                <p><strong>Преподаватель:</strong> ${subject.teacher}</p>
            </div>
        `).join('');
    }
}