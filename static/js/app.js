const API_URL = 'http://localhost:8080/api/v1';

// Check if user is logged in
function checkAuth() {
    const token = localStorage.getItem('token');
    const username = localStorage.getItem('username');
    
    if (token && username) {
        document.getElementById('login-link')?.classList.add('hidden');
        document.getElementById('register-link')?.classList.add('hidden');
        document.getElementById('logout-btn')?.classList.remove('hidden');
        
        const welcomeMsg = document.getElementById('welcome-message');
        if (welcomeMsg) {
            welcomeMsg.textContent = `Welcome back, ${username}! Ready to cook?`;
        }
    }
}

// Login form handler
const loginForm = document.getElementById('login-form');
if (loginForm) {
    loginForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;
        const errorDiv = document.getElementById('error-message');
        const successDiv = document.getElementById('success-message');
        
        errorDiv.classList.add('hidden');
        successDiv.classList.add('hidden');
        
        try {
            const response = await fetch(`${API_URL}/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `username=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`
            });
            
            if (response.ok) {
                const data = await response.json();
                localStorage.setItem('token', data.token);
                localStorage.setItem('username', username);
                
                successDiv.textContent = 'Login successful! Redirecting...';
                successDiv.classList.remove('hidden');
                
                setTimeout(() => {
                    window.location.href = '/';
                }, 1000);
            } else {
                const errorText = await response.text();
                errorDiv.textContent = errorText || 'Login failed. Please check your credentials.';
                errorDiv.classList.remove('hidden');
            }
        } catch (error) {
            errorDiv.textContent = 'Network error. Please try again.';
            errorDiv.classList.remove('hidden');
        }
    });
}

// Register form handler
const registerForm = document.getElementById('register-form');
if (registerForm) {
    registerForm.addEventListener('submit', async (e) => {
        e.preventDefault();
        
        const username = document.getElementById('reg-username').value;
        const password = document.getElementById('reg-password').value;
        const errorDiv = document.getElementById('reg-error-message');
        const successDiv = document.getElementById('reg-success-message');
        
        errorDiv.classList.add('hidden');
        successDiv.classList.add('hidden');
        
        try {
            const response = await fetch(`${API_URL}/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
                body: `username=${encodeURIComponent(username)}&password=${encodeURIComponent(password)}`
            });
            
            if (response.ok) {
                successDiv.textContent = 'Registration successful! Redirecting to login...';
                successDiv.classList.remove('hidden');
                
                setTimeout(() => {
                    window.location.href = '/login';
                }, 1500);
            } else {
                const errorText = await response.text();
                errorDiv.textContent = errorText || 'Registration failed. Username might already exist.';
                errorDiv.classList.remove('hidden');
            }
        } catch (error) {
            errorDiv.textContent = 'Network error. Please try again.';
            errorDiv.classList.remove('hidden');
        }
    });
}

// Logout function
function logout() {
    fetch(`${API_URL}/logout`, {
        method: 'POST',
        headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        }
    }).finally(() => {
        localStorage.removeItem('token');
        localStorage.removeItem('username');
        window.location.href = '/';
    });
}

// Check auth on page load
document.addEventListener('DOMContentLoaded', checkAuth);