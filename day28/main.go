/*package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi! welcome to season 10 tutorial")
}

func main() {
	http.HandleFunc("/", helloHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Home Page!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm Amirali and I'm a student")
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/about", aboutPage)

	fmt.Println("Server is running on http://localhost:9090")
	http.ListenAndServe(":9090", nil)

}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func methodHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "This is a GET request")
	} else if r.Method == http.MethodPost {
		fmt.Fprintf(w, "This is a POST request")
	} else {
		fmt.Fprintf(w, "Other method used: %s", r.Method)
	}
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	userID := queryParams.Get("id")

	if userID == "" {
		fmt.Fprintf(w, "No ID provided in the URL")
	} else {
		fmt.Fprintf(w, "User ID requested: %s", userID)
	}
}

func main() {
	http.HandleFunc("/check-method", methodHandler)

	http.HandleFunc("/user", queryHandler)

	fmt.Println("Server starting on http://localhost:9090")

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
*/

/*
	package main

	import (
		"fmt"
		"net/http"
	)

	func greeting(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Home Page")
	}

	func about(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I'm Amirali and I'm a student")
	}

	func hello(w http.ResponseWriter, r *http.Request) {
		queryDecode := r.URL.Query()

		name := queryDecode.Get("name")

		if name == "" {
			fmt.Fprintf(w, "Hello Guest")
		} else {
			fmt.Fprintf(w, "Hello %s, Welcome to Go course", name)
		}
	}

	func main() {
		http.HandleFunc("/", greeting)
		http.HandleFunc("/about", about)
		http.HandleFunc("/hello", hello)

		fmt.Println("Server starting on http://localhost:9090")

		err := http.ListenAndServe(":9090", nil)
		if err != nil {
			fmt.Println("Error starting server:", err)
		}

	}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./assets"))

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome! To see a file, go to /static/ID.png")
	})

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func staticFile(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Use /files/ path to access static content")
}

func main() {
	fileServer := http.FileServer(http.Dir("./public"))
	http.Handle("/files/", http.StripPrefix("/files/", fileServer))
	http.HandleFunc("/", staticFile)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>ورود به حساب کاربری</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Vazirmatn:wght@300;400;500;600;700&display=swap" rel="stylesheet">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Vazirmatn', sans-serif;
        }

        body {
            background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
            overflow-x: hidden;
        }

        .container {
            width: 100%;
            max-width: 450px;
            perspective: 1000px;
        }

        .login-card {
            background-color: rgba(255, 255, 255, 0.95);
            border-radius: 25px;
            padding: 40px 35px;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.1);
            transform-style: preserve-3d;
            transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            position: relative;
            overflow: hidden;
        }

        .login-card:hover {
            transform: rotateY(5deg) rotateX(2deg) scale(1.02);
            box-shadow: 0 30px 60px rgba(0, 0, 0, 0.3), 0 0 0 1px rgba(255, 255, 255, 0.15);
        }

        .login-card:before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 5px;
            background: linear-gradient(90deg, #6a11cb, #2575fc, #6a11cb);
            background-size: 200% 100%;
            animation: shimmer 3s infinite linear;
        }

        .header {
            text-align: center;
            margin-bottom: 40px;
            position: relative;
        }

        .logo {
            width: 80px;
            height: 80px;
            background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
            border-radius: 50%;
            margin: 0 auto 20px;
            display: flex;
            align-items: center;
            justify-content: center;
            box-shadow: 0 10px 20px rgba(106, 17, 203, 0.3);
            animation: float 3s ease-in-out infinite;
        }

        .logo i {
            font-size: 36px;
            color: white;
        }

        .header h1 {
            font-size: 28px;
            color: #333;
            font-weight: 700;
            margin-bottom: 8px;
        }

        .header p {
            color: #666;
            font-size: 15px;
        }

        .input-group {
            margin-bottom: 25px;
            position: relative;
        }

        .input-group label {
            display: block;
            margin-bottom: 8px;
            color: #555;
            font-weight: 500;
            font-size: 15px;
            transition: color 0.3s;
        }

        .input-field {
            width: 100%;
            padding: 15px 20px;
            border: 2px solid #e1e1e1;
            border-radius: 15px;
            font-size: 16px;
            transition: all 0.3s;
            background-color: #f9f9f9;
            color: #333;
        }

        .input-field:focus {
            outline: none;
            border-color: #6a11cb;
            box-shadow: 0 0 0 3px rgba(106, 17, 203, 0.1);
            background-color: white;
            transform: translateY(-2px);
        }

        .input-field:focus + .input-icon {
            color: #6a11cb;
            transform: scale(1.1);
        }

        .input-icon {
            position: absolute;
            left: 20px;
            top: 43px;
            color: #999;
            font-size: 18px;
            transition: all 0.3s;
        }

        .options {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-bottom: 30px;
            font-size: 14px;
        }

        .remember-me {
            display: flex;
            align-items: center;
        }

        .remember-me input {
            margin-left: 8px;
            width: 18px;
            height: 18px;
            accent-color: #6a11cb;
        }

        .forgot-password {
            color: #6a11cb;
            text-decoration: none;
            font-weight: 500;
            transition: color 0.3s;
            position: relative;
        }

        .forgot-password:after {
            content: '';
            position: absolute;
            bottom: -2px;
            right: 0;
            width: 0;
            height: 2px;
            background-color: #6a11cb;
            transition: width 0.3s;
        }

        .forgot-password:hover:after {
            width: 100%;
        }

        .forgot-password:hover {
            color: #2575fc;
        }

        .login-btn {
            width: 100%;
            padding: 17px;
            background: linear-gradient(135deg, #6a11cb 0%, #2575fc 100%);
            color: white;
            border: none;
            border-radius: 15px;
            font-size: 17px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s;
            box-shadow: 0 10px 20px rgba(106, 17, 203, 0.3);
            position: relative;
            overflow: hidden;
        }

        .login-btn:hover {
            transform: translateY(-3px);
            box-shadow: 0 15px 25px rgba(106, 17, 203, 0.4);
        }

        .login-btn:active {
            transform: translateY(0);
        }

        .login-btn:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
            transition: left 0.7s;
        }

        .login-btn:hover:before {
            left: 100%;
        }

        .divider {
            text-align: center;
            margin: 30px 0;
            position: relative;
            color: #999;
            font-size: 14px;
        }

        .divider:before {
            content: '';
            position: absolute;
            top: 50%;
            right: 0;
            width: 42%;
            height: 1px;
            background-color: #e1e1e1;
        }

        .divider:after {
            content: '';
            position: absolute;
            top: 50%;
            left: 0;
            width: 42%;
            height: 1px;
            background-color: #e1e1e1;
        }

        .social-login {
            display: flex;
            justify-content: center;
            gap: 20px;
        }

        .social-btn {
            width: 55px;
            height: 55px;
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            font-size: 22px;
            color: white;
            cursor: pointer;
            transition: all 0.3s;
            box-shadow: 0 8px 15px rgba(0, 0, 0, 0.1);
        }

        .social-btn:hover {
            transform: translateY(-5px);
            box-shadow: 0 12px 20px rgba(0, 0, 0, 0.15);
        }

        .google {
            background: linear-gradient(135deg, #dd4b39, #c23321);
        }

        .facebook {
            background: linear-gradient(135deg, #3b5998, #2d4373);
        }

        .twitter {
            background: linear-gradient(135deg, #1da1f2, #0d8bd9);
        }

        .register {
            text-align: center;
            margin-top: 30px;
            color: #666;
            font-size: 15px;
        }

        .register a {
            color: #6a11cb;
            text-decoration: none;
            font-weight: 600;
            margin-right: 5px;
            transition: color 0.3s;
        }

        .register a:hover {
            color: #2575fc;
            text-decoration: underline;
        }

        .floating-balls {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            overflow: hidden;
            z-index: -1;
            border-radius: 25px;
        }

        .ball {
            position: absolute;
            border-radius: 50%;
            background: rgba(106, 17, 203, 0.1);
            animation: float-ball 15s infinite linear;
        }

        .ball:nth-child(1) {
            width: 80px;
            height: 80px;
            top: 10%;
            left: 10%;
            animation-delay: 0s;
        }

        .ball:nth-child(2) {
            width: 60px;
            height: 60px;
            top: 70%;
            left: 80%;
            animation-delay: -3s;
        }

        .ball:nth-child(3) {
            width: 100px;
            height: 100px;
            top: 50%;
            left: 20%;
            animation-delay: -6s;
        }

        .ball:nth-child(4) {
            width: 50px;
            height: 50px;
            top: 20%;
            left: 85%;
            animation-delay: -9s;
        }

        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-15px);
            }
        }

        @keyframes float-ball {
            0% {
                transform: translateY(0) rotate(0deg);
                opacity: 0.7;
            }
            33% {
                transform: translateY(-30px) rotate(120deg);
                opacity: 0.4;
            }
            66% {
                transform: translateY(20px) rotate(240deg);
                opacity: 0.8;
            }
            100% {
                transform: translateY(0) rotate(360deg);
                opacity: 0.7;
            }
        }

        @keyframes shimmer {
            0% {
                background-position: 200% 0;
            }
            100% {
                background-position: -200% 0;
            }
        }

        @keyframes shake {
            0%, 100% { transform: translateX(0); }
            10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
            20%, 40%, 60%, 80% { transform: translateX(5px); }
        }

        .shake {
            animation: shake 0.5s ease-in-out;
        }

        .success-message {
            position: fixed;
            top: 30px;
            right: 30px;
            background: linear-gradient(135deg, #4CAF50, #2E7D32);
            color: white;
            padding: 15px 25px;
            border-radius: 10px;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
            display: flex;
            align-items: center;
            gap: 15px;
            transform: translateX(150%);
            transition: transform 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
            z-index: 1000;
        }

        .success-message.show {
            transform: translateX(0);
        }

        .success-message i {
            font-size: 24px;
        }

        @keyframes card-entry {
            0% {
                opacity: 0;
                transform: translateY(30px) rotateX(30deg);
            }
            100% {
                opacity: 1;
                transform: translateY(0) rotateX(0);
            }
        }

        .login-card {
            animation: card-entry 0.8s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
        }

        @media (max-width: 500px) {
            .login-card {
                padding: 30px 25px;
            }

            .header h1 {
                font-size: 24px;
            }

            .options {
                flex-direction: column;
                align-items: flex-start;
                gap: 15px;
            }

            .social-login {
                gap: 15px;
            }

            .social-btn {
                width: 50px;
                height: 50px;
                font-size: 20px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="login-card">
            <div class="floating-balls">
                <div class="ball"></div>
                <div class="ball"></div>
                <div class="ball"></div>
                <div class="ball"></div>
            </div>

            <div class="header">
                <div class="logo">
                    <i class="fas fa-user-lock"></i>
                </div>
                <h1>ورود به حساب کاربری</h1>
                <p>لطفا اطلاعات حساب خود را وارد کنید</p>
            </div>

            <form id="loginForm">
                <div class="input-group">
                    <label for="username"><i class="fas fa-user"></i> نام کاربری</label>
                    <input type="text" id="username" class="input-field" placeholder="نام کاربری خود را وارد کنید" required>
                    <div class="input-icon">
                        <i class="fas fa-user"></i>
                    </div>
                </div>

                <div class="input-group">
                    <label for="password"><i class="fas fa-lock"></i> رمز عبور</label>
                    <input type="password" id="password" class="input-field" placeholder="رمز عبور خود را وارد کنید" required>
                    <div class="input-icon">
                        <i class="fas fa-lock"></i>
                    </div>
                </div>

                <div class="options">
                    <div class="remember-me">
                        <input type="checkbox" id="remember">
                        <label for="remember">مرا به خاطر بسپار</label>
                    </div>
                    <a href="#" class="forgot-password">رمز عبور را فراموش کرده‌اید؟</a>
                </div>

                <button type="submit" class="login-btn">
                    <i class="fas fa-sign-in-alt"></i> ورود به حساب
                </button>
            </form>

            <div class="divider">
                یا وارد شوید با
            </div>

            <div class="social-login">
                <div class="social-btn google">
                    <i class="fab fa-google"></i>
                </div>
                <div class="social-btn facebook">
                    <i class="fab fa-facebook-f"></i>
                </div>
                <div class="social-btn twitter">
                    <i class="fab fa-twitter"></i>
                </div>
            </div>

            <div class="register">
                حساب کاربری ندارید؟ <a href="#">ثبت‌نام کنید</a>
            </div>
        </div>
    </div>

    <div class="success-message" id="successMessage">
        <i class="fas fa-check-circle"></i>
        <div>ورود با موفقیت انجام شد!</div>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const loginForm = document.getElementById('loginForm');
            const loginCard = document.querySelector('.login-card');
            const successMessage = document.getElementById('successMessage');
            const socialButtons = document.querySelectorAll('.social-btn');

            setTimeout(() => {
                loginCard.style.opacity = '1';
            }, 100);

            loginForm.addEventListener('submit', function(e) {
                e.preventDefault();

                const username = document.getElementById('username').value;
                const password = document.getElementById('password').value;

                if (username.length < 3 || password.length < 6) {
                    loginCard.classList.add('shake');
                    setTimeout(() => {
                        loginCard.classList.remove('shake');
                    }, 500);

                    const inputs = document.querySelectorAll('.input-field');
                    inputs.forEach(input => {
                        if (input.value.length === 0) {
                            input.style.borderColor = '#ff4757';
                            setTimeout(() => {
                                input.style.borderColor = '#e1e1e1';
                            }, 2000);
                        }
                    });

                    return;
                }

                loginCard.style.transform = 'scale(0.98)';
                setTimeout(() => {
                    loginCard.style.transform = 'scale(1)';
                    successMessage.classList.add('show');

                    loginForm.reset();

                    setTimeout(() => {
                        successMessage.classList.remove('show');
                    }, 3000);
                }, 300);
            });

            socialButtons.forEach(btn => {
                btn.addEventListener('click', function() {
                    this.style.transform = 'scale(0.9)';
                    setTimeout(() => {
                        this.style.transform = 'scale(1) translateY(-5px)';
                    }, 150);

                    const platform = this.classList.contains('google') ? 'گوگل' :
                                   this.classList.contains('facebook') ? 'فیس‌بوک' : 'توییتر';

                    successMessage.innerHTML = '<i class="fas fa-info-circle"></i><div>ورود با ' + platform + ' در حال توسعه است</div>';

                    successMessage.style.background = 'linear-gradient(135deg, #2575fc, #6a11cb)';
                    successMessage.classList.add('show');

                    setTimeout(() => {
                        successMessage.classList.remove('show');
                    }, 2500);
                });
            });

            const inputFields = document.querySelectorAll('.input-field');
            inputFields.forEach(input => {
                input.addEventListener('focus', function() {
                    this.parentElement.querySelector('label').style.color = '#6a11cb';
                });

                input.addEventListener('blur', function() {
                    this.parentElement.querySelector('label').style.color = '#555';
                });
            });

            const forgotLink = document.querySelector('.forgot-password');
            forgotLink.addEventListener('click', function(e) {
                e.preventDefault();

                const passwordField = document.getElementById('password');
                passwordField.style.borderColor = '#ffa502';
                passwordField.style.boxShadow = '0 0 0 3px rgba(255, 165, 2, 0.2)';

                setTimeout(() => {
                    passwordField.style.borderColor = '#e1e1e1';
                    passwordField.style.boxShadow = 'none';
                }, 1500);

                successMessage.innerHTML = '<i class="fas fa-envelope"></i><div>لینک بازنشانی رمز عبور به ایمیل شما ارسال شد</div>';
                successMessage.style.background = 'linear-gradient(135deg, #ffa502, #ff7f00)';
                successMessage.classList.add('show');

                setTimeout(() => {
                    successMessage.classList.remove('show');
                }, 3000);
            });
        });
    </script>
</body>
</html>`
		fmt.Fprint(w, html)
	} else if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintf(w, "Error parsing form: %v", err)
			return
		}
		user := r.FormValue("username")
		pass := r.FormValue("password")

		fmt.Fprintf(w, "Received - Username: %s, Password: %s", user, pass)
	}
}

func main() {
	http.HandleFunc("/login", loginHandler)

	fmt.Println("Server starting on http://localhost:8080/login")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"net/http"
)

func contact(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `
		<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>فرم تماس با ما</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Vazirmatn:wght@300;400;500;600;700&display=swap" rel="stylesheet">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Vazirmatn', sans-serif;
        }

        body {
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
            overflow-x: hidden;
        }

        .container {
            width: 100%;
            max-width: 500px;
            perspective: 1000px;
        }

        .contact-card {
            background-color: rgba(255, 255, 255, 0.95);
            border-radius: 25px;
            padding: 40px 35px;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.15), 0 0 0 1px rgba(255, 255, 255, 0.1);
            transform-style: preserve-3d;
            transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            position: relative;
            overflow: hidden;
        }

        .contact-card:hover {
            transform: rotateY(3deg) rotateX(1deg) scale(1.01);
            box-shadow: 0 25px 50px rgba(0, 0, 0, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.15);
        }

        .contact-card:before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 5px;
            background: linear-gradient(90deg, #11998e, #38ef7d, #11998e);
            background-size: 200% 100%;
            animation: shimmer 3s infinite linear;
        }

        .header {
            text-align: center;
            margin-bottom: 40px;
            position: relative;
        }

        .logo {
            width: 80px;
            height: 80px;
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            border-radius: 50%;
            margin: 0 auto 20px;
            display: flex;
            align-items: center;
            justify-content: center;
            box-shadow: 0 10px 20px rgba(17, 153, 142, 0.3);
            animation: float 3s ease-in-out infinite;
        }

        .logo i {
            font-size: 36px;
            color: white;
        }

        .header h1 {
            font-size: 28px;
            color: #333;
            font-weight: 700;
            margin-bottom: 8px;
        }

        .header p {
            color: #666;
            font-size: 15px;
            line-height: 1.6;
        }

        .input-group {
            margin-bottom: 25px;
            position: relative;
        }

        .input-group label {
            display: block;
            margin-bottom: 8px;
            color: #555;
            font-weight: 500;
            font-size: 15px;
            transition: color 0.3s;
        }

        .input-field {
            width: 100%;
            padding: 15px 20px;
            border: 2px solid #e1e1e1;
            border-radius: 15px;
            font-size: 16px;
            transition: all 0.3s;
            background-color: #f9f9f9;
            color: #333;
        }

        .input-field:focus {
            outline: none;
            border-color: #11998e;
            box-shadow: 0 0 0 3px rgba(17, 153, 142, 0.1);
            background-color: white;
            transform: translateY(-2px);
        }

        .input-field:focus + .input-icon {
            color: #11998e;
            transform: scale(1.1);
        }

        .input-icon {
            position: absolute;
            left: 20px;
            top: 43px;
            color: #999;
            font-size: 18px;
            transition: all 0.3s;
        }

        textarea.input-field {
            min-height: 150px;
            resize: vertical;
            line-height: 1.6;
        }

        .form-actions {
            display: flex;
            justify-content: space-between;
            align-items: center;
            margin-top: 30px;
        }

        .submit-btn {
            padding: 17px 35px;
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            color: white;
            border: none;
            border-radius: 15px;
            font-size: 17px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s;
            box-shadow: 0 10px 20px rgba(17, 153, 142, 0.3);
            position: relative;
            overflow: hidden;
            display: flex;
            align-items: center;
            gap: 10px;
        }

        .submit-btn:hover {
            transform: translateY(-3px);
            box-shadow: 0 15px 25px rgba(17, 153, 142, 0.4);
        }

        .submit-btn:active {
            transform: translateY(0);
        }

        .submit-btn:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
            transition: left 0.7s;
        }

        .submit-btn:hover:before {
            left: 100%;
        }

        .reset-btn {
            padding: 15px 25px;
            background: #f5f5f5;
            color: #666;
            border: 2px solid #e1e1e1;
            border-radius: 15px;
            font-size: 16px;
            font-weight: 500;
            cursor: pointer;
            transition: all 0.3s;
        }

        .reset-btn:hover {
            background: #e9e9e9;
            transform: translateY(-2px);
        }

        .contact-info {
            margin-top: 40px;
            padding-top: 25px;
            border-top: 1px solid #e1e1e1;
            display: flex;
            justify-content: space-around;
            text-align: center;
        }

        .info-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 8px;
        }

        .info-icon {
            width: 45px;
            height: 45px;
            background: linear-gradient(135deg, #11998e 0%, #38ef7d 100%);
            border-radius: 50%;
            display: flex;
            align-items: center;
            justify-content: center;
            color: white;
            font-size: 18px;
        }

        .info-text {
            color: #666;
            font-size: 14px;
        }

        .floating-elements {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            overflow: hidden;
            z-index: -1;
            border-radius: 25px;
        }

        .element {
            position: absolute;
            background: rgba(17, 153, 142, 0.08);
            animation: float-element 20s infinite linear;
        }

        .element:nth-child(1) {
            width: 60px;
            height: 60px;
            top: 10%;
            left: 10%;
            border-radius: 50%;
            animation-delay: 0s;
        }

        .element:nth-child(2) {
            width: 40px;
            height: 40px;
            top: 70%;
            left: 80%;
            border-radius: 10px;
            animation-delay: -5s;
        }

        .element:nth-child(3) {
            width: 80px;
            height: 80px;
            top: 50%;
            left: 20%;
            border-radius: 50%;
            animation-delay: -10s;
        }

        .element:nth-child(4) {
            width: 50px;
            height: 50px;
            top: 20%;
            left: 85%;
            border-radius: 15px;
            animation-delay: -15s;
        }

        .success-message {
            position: fixed;
            top: 30px;
            right: 30px;
            background: linear-gradient(135deg, #11998e, #38ef7d);
            color: white;
            padding: 15px 25px;
            border-radius: 10px;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
            display: flex;
            align-items: center;
            gap: 15px;
            transform: translateX(150%);
            transition: transform 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
            z-index: 1000;
        }

        .success-message.show {
            transform: translateX(0);
        }

        .success-message i {
            font-size: 24px;
        }

        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-15px);
            }
        }

        @keyframes float-element {
            0% {
                transform: translateY(0) rotate(0deg);
                opacity: 0.7;
            }
            33% {
                transform: translateY(-30px) rotate(120deg);
                opacity: 0.4;
            }
            66% {
                transform: translateY(20px) rotate(240deg);
                opacity: 0.8;
            }
            100% {
                transform: translateY(0) rotate(360deg);
                opacity: 0.7;
            }
        }

        @keyframes shimmer {
            0% {
                background-position: 200% 0;
            }
            100% {
                background-position: -200% 0;
            }
        }

        @keyframes shake {
            0%, 100% { transform: translateX(0); }
            10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
            20%, 40%, 60%, 80% { transform: translateX(5px); }
        }

        .shake {
            animation: shake 0.5s ease-in-out;
        }

        @keyframes card-entry {
            0% {
                opacity: 0;
                transform: translateY(30px) rotateX(30deg);
            }
            100% {
                opacity: 1;
                transform: translateY(0) rotateX(0);
            }
        }

        .contact-card {
            animation: card-entry 0.8s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
        }

        @media (max-width: 600px) {
            .contact-card {
                padding: 30px 25px;
            }

            .header h1 {
                font-size: 24px;
            }

            .form-actions {
                flex-direction: column;
                gap: 15px;
            }

            .submit-btn, .reset-btn {
                width: 100%;
                justify-content: center;
            }

            .contact-info {
                flex-direction: column;
                gap: 20px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="contact-card">
            <div class="floating-elements">
                <div class="element"></div>
                <div class="element"></div>
                <div class="element"></div>
                <div class="element"></div>
            </div>

            <div class="header">
                <div class="logo">
                    <i class="fas fa-envelope-open-text"></i>
                </div>
                <h1>تماس با ما</h1>
                <p>نظرات، پیشنهادات و انتقادات خود را با ما در میان بگذارید. ما در اسرع وقت پاسخگوی شما خواهیم بود.</p>
            </div>

            <form method="POST" action="/contact" id="contactForm">
                <div class="input-group">
                    <label for="email"><i class="fas fa-envelope"></i> آدرس ایمیل</label>
                    <input type="email" id="email" name="email" class="input-field" placeholder="example@domain.com" required>
                    <div class="input-icon">
                        <i class="fas fa-envelope"></i>
                    </div>
                </div>

                <div class="input-group">
                    <label for="message"><i class="fas fa-comment-dots"></i> پیام شما</label>
                    <textarea id="message" name="message" class="input-field" placeholder="پیام خود را در اینجا بنویسید..." required></textarea>
                    <div class="input-icon">
                        <i class="fas fa-comment-dots"></i>
                    </div>
                </div>

                <div class="form-actions">
                    <button type="submit" class="submit-btn">
                        <i class="fas fa-paper-plane"></i> ارسال پیام
                    </button>
                    <button type="reset" class="reset-btn">
                        <i class="fas fa-redo"></i> بازنشانی فرم
                    </button>
                </div>
            </form>

            <div class="contact-info">
                <div class="info-item">
                    <div class="info-icon">
                        <i class="fas fa-clock"></i>
                    </div>
                    <div class="info-text">پاسخگویی ۲۴ ساعته</div>
                </div>

                <div class="info-item">
                    <div class="info-icon">
                        <i class="fas fa-headset"></i>
                    </div>
                    <div class="info-text">پشتیبانی آنلاین</div>
                </div>

                <div class="info-item">
                    <div class="info-icon">
                        <i class="fas fa-shield-alt"></i>
                    </div>
                    <div class="info-text">حریم خصوصی شما</div>
                </div>
            </div>
        </div>
    </div>

    <div class="success-message" id="successMessage">
        <i class="fas fa-check-circle"></i>
        <div>پیام شما با موفقیت ارسال شد!</div>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const contactForm = document.getElementById('contactForm');
            const contactCard = document.querySelector('.contact-card');
            const successMessage = document.getElementById('successMessage');

            // انیمیشن ورود اولیه
            setTimeout(() => {
                contactCard.style.opacity = '1';
            }, 100);

            // اعتبارسنجی و ارسال فرم
            contactForm.addEventListener('submit', function(e) {
                e.preventDefault();

                const email = document.getElementById('email').value;
                const message = document.getElementById('message').value;

                // اعتبارسنجی ساده
                const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

                if (!emailPattern.test(email)) {
                    // انیمیشن لرزش برای نشان دادن خطا
                    contactCard.classList.add('shake');
                    setTimeout(() => {
                        contactCard.classList.remove('shake');
                    }, 500);

                    // تغییر رنگ border فیلد ایمیل
                    const emailField = document.getElementById('email');
                    emailField.style.borderColor = '#ff4757';
                    emailField.style.boxShadow = '0 0 0 3px rgba(255, 71, 87, 0.1)';

                    setTimeout(() => {
                        emailField.style.borderColor = '#e1e1e1';
                        emailField.style.boxShadow = 'none';
                    }, 2000);

                    return;
                }

                if (message.length < 10) {
                    // انیمیشن لرزش برای نشان دادن خطا
                    contactCard.classList.add('shake');
                    setTimeout(() => {
                        contactCard.classList.remove('shake');
                    }, 500);

                    // تغییر رنگ border فیلد پیام
                    const messageField = document.getElementById('message');
                    messageField.style.borderColor = '#ff4757';
                    messageField.style.boxShadow = '0 0 0 3px rgba(255, 71, 87, 0.1)';

                    setTimeout(() => {
                        messageField.style.borderColor = '#e1e1e1';
                        messageField.style.boxShadow = 'none';
                    }, 2000);

                    return;
                }

                // شبیه‌سازی ارسال موفق
                contactCard.style.transform = 'scale(0.98)';

                // نمایش انیمیشن ارسال
                const submitBtn = contactForm.querySelector('.submit-btn');
                const originalText = submitBtn.innerHTML;
                submitBtn.innerHTML = '<i class="fas fa-spinner fa-spin"></i> در حال ارسال...';
                submitBtn.disabled = true;

                // در اینجا فرم واقعی ارسال می‌شود
                setTimeout(() => {
                    // ارسال واقعی فرم (جاوااسکریپت submit می‌کند)
                    contactForm.submit();
                }, 1500);
            });

            // انیمیشن focus برای فیلدهای ورودی
            const inputFields = document.querySelectorAll('.input-field');
            inputFields.forEach(input => {
                input.addEventListener('focus', function() {
                    this.parentElement.querySelector('label').style.color = '#11998e';
                });

                input.addEventListener('blur', function() {
                    this.parentElement.querySelector('label').style.color = '#555';
                });
            });

            // دکمه بازنشانی
            const resetBtn = contactForm.querySelector('button[type="reset"]');
            resetBtn.addEventListener('click', function() {
                // انیمیشن کوچک شدن و برگشت
                contactCard.style.transform = 'scale(0.95)';
                setTimeout(() => {
                    contactCard.style.transform = 'scale(1)';
                }, 300);

                // نمایش پیام
                successMessage.innerHTML = '<i class="fas fa-undo"></i><div>فرم با موفقیت بازنشانی شد</div>';
                successMessage.style.background = 'linear-gradient(135deg, #ffa502, #ff7f00)';
                successMessage.classList.add('show');

                setTimeout(() => {
                    successMessage.classList.remove('show');
                }, 2500);
            });
        });
    </script>
</body>
</html>
		`
		fmt.Fprint(w, html)
		return
	}

	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		message := r.FormValue("message")

		fmt.Fprintf(w, "Message sent from %s: %s", email, message)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/contact", contact)

	fmt.Println("Server starting on http://localhost:8080/contact")
	http.ListenAndServe(":8080", nil)
}
*/

/*
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `
		<!DOCTYPE html>
<html lang="fa" dir="rtl">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>آپلود فایل</title>
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.0/css/all.min.css">
    <link href="https://fonts.googleapis.com/css2?family=Vazirmatn:wght@300;400;500;600;700&display=swap" rel="stylesheet">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: 'Vazirmatn', sans-serif;
        }

        body {
            background: linear-gradient(135deg, #8A2BE2 0%, #4B0082 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
            overflow-x: hidden;
        }

        .container {
            width: 100%;
            max-width: 600px;
            perspective: 1000px;
        }

        .upload-card {
            background-color: rgba(255, 255, 255, 0.95);
            border-radius: 25px;
            padding: 40px 35px;
            box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2), 0 0 0 1px rgba(255, 255, 255, 0.1);
            transform-style: preserve-3d;
            transition: all 0.5s cubic-bezier(0.175, 0.885, 0.32, 1.275);
            position: relative;
            overflow: hidden;
        }

        .upload-card:hover {
            transform: rotateY(3deg) rotateX(1deg) scale(1.01);
            box-shadow: 0 25px 50px rgba(0, 0, 0, 0.25), 0 0 0 1px rgba(255, 255, 255, 0.15);
        }

        .upload-card:before {
            content: '';
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 5px;
            background: linear-gradient(90deg, #8A2BE2, #4B0082, #8A2BE2);
            background-size: 200% 100%;
            animation: shimmer 3s infinite linear;
        }

        .header {
            text-align: center;
            margin-bottom: 40px;
            position: relative;
        }

        .logo {
            width: 80px;
            height: 80px;
            background: linear-gradient(135deg, #8A2BE2 0%, #4B0082 100%);
            border-radius: 50%;
            margin: 0 auto 20px;
            display: flex;
            align-items: center;
            justify-content: center;
            box-shadow: 0 10px 20px rgba(138, 43, 226, 0.3);
            animation: float 3s ease-in-out infinite;
        }

        .logo i {
            font-size: 36px;
            color: white;
        }

        .header h1 {
            font-size: 28px;
            color: #333;
            font-weight: 700;
            margin-bottom: 8px;
        }

        .header p {
            color: #666;
            font-size: 15px;
            line-height: 1.6;
        }

        .upload-area {
            border: 3px dashed #8A2BE2;
            border-radius: 20px;
            padding: 40px 20px;
            text-align: center;
            background-color: rgba(138, 43, 226, 0.05);
            margin-bottom: 30px;
            cursor: pointer;
            transition: all 0.3s;
            position: relative;
        }

        .upload-area:hover {
            background-color: rgba(138, 43, 226, 0.1);
            border-color: #4B0082;
            transform: translateY(-3px);
        }

        .upload-area.dragover {
            background-color: rgba(138, 43, 226, 0.15);
            border-color: #4B0082;
            transform: scale(1.02);
        }

        .upload-icon {
            font-size: 60px;
            color: #8A2BE2;
            margin-bottom: 20px;
        }

        .upload-text {
            color: #555;
            font-size: 18px;
            margin-bottom: 10px;
        }

        .upload-subtext {
            color: #888;
            font-size: 14px;
        }

        .file-input {
            display: none;
        }

        .browse-btn {
            display: inline-block;
            margin-top: 15px;
            padding: 10px 25px;
            background: linear-gradient(135deg, #8A2BE2 0%, #4B0082 100%);
            color: white;
            border: none;
            border-radius: 10px;
            font-size: 16px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s;
            box-shadow: 0 5px 15px rgba(138, 43, 226, 0.3);
        }

        .browse-btn:hover {
            transform: translateY(-2px);
            box-shadow: 0 8px 20px rgba(138, 43, 226, 0.4);
        }

        .file-info {
            background: #f8f9fa;
            border-radius: 15px;
            padding: 20px;
            margin-top: 20px;
            display: none;
            animation: fadeIn 0.5s ease-out;
        }

        .file-info.show {
            display: block;
        }

        .file-preview {
            margin-bottom: 20px;
            text-align: center;
        }

        .preview-image {
            max-width: 100%;
            max-height: 200px;
            border-radius: 10px;
            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
        }

        .preview-icon {
            font-size: 80px;
            color: #8A2BE2;
        }

        .file-details {
            text-align: right;
        }

        .file-detail {
            margin-bottom: 10px;
            display: flex;
            justify-content: space-between;
            padding-bottom: 8px;
            border-bottom: 1px solid #eee;
        }

        .detail-label {
            color: #666;
            font-weight: 500;
        }

        .detail-value {
            color: #333;
            font-weight: 600;
        }

        .submit-btn {
            width: 100%;
            padding: 17px;
            background: linear-gradient(135deg, #8A2BE2 0%, #4B0082 100%);
            color: white;
            border: none;
            border-radius: 15px;
            font-size: 17px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.3s;
            box-shadow: 0 10px 20px rgba(138, 43, 226, 0.3);
            position: relative;
            overflow: hidden;
            display: flex;
            align-items: center;
            justify-content: center;
            gap: 10px;
            margin-top: 20px;
        }

        .submit-btn:hover:not(:disabled) {
            transform: translateY(-3px);
            box-shadow: 0 15px 25px rgba(138, 43, 226, 0.4);
        }

        .submit-btn:disabled {
            opacity: 0.6;
            cursor: not-allowed;
        }

        .submit-btn:before {
            content: '';
            position: absolute;
            top: 0;
            left: -100%;
            width: 100%;
            height: 100%;
            background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.2), transparent);
            transition: left 0.7s;
        }

        .submit-btn:hover:not(:disabled):before {
            left: 100%;
        }

        .floating-elements {
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            overflow: hidden;
            z-index: -1;
            border-radius: 25px;
        }

        .element {
            position: absolute;
            background: rgba(138, 43, 226, 0.08);
            animation: float-element 20s infinite linear;
        }

        .element:nth-child(1) {
            width: 60px;
            height: 60px;
            top: 10%;
            left: 10%;
            border-radius: 50%;
            animation-delay: 0s;
        }

        .element:nth-child(2) {
            width: 40px;
            height: 40px;
            top: 70%;
            left: 80%;
            border-radius: 10px;
            animation-delay: -5s;
        }

        .element:nth-child(3) {
            width: 80px;
            height: 80px;
            top: 50%;
            left: 20%;
            border-radius: 50%;
            animation-delay: -10s;
        }

        .element:nth-child(4) {
            width: 50px;
            height: 50px;
            top: 20%;
            left: 85%;
            border-radius: 15px;
            animation-delay: -15s;
        }

        .progress-bar {
            width: 100%;
            height: 10px;
            background-color: #e9ecef;
            border-radius: 5px;
            margin-top: 15px;
            overflow: hidden;
            display: none;
        }

        .progress-fill {
            height: 100%;
            background: linear-gradient(90deg, #8A2BE2, #4B0082);
            width: 0%;
            border-radius: 5px;
            transition: width 0.3s ease;
        }

        .success-message {
            position: fixed;
            top: 30px;
            right: 30px;
            background: linear-gradient(135deg, #38ef7d, #11998e);
            color: white;
            padding: 15px 25px;
            border-radius: 10px;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
            display: flex;
            align-items: center;
            gap: 15px;
            transform: translateX(150%);
            transition: transform 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
            z-index: 1000;
        }

        .success-message.show {
            transform: translateX(0);
        }

        .error-message {
            position: fixed;
            top: 30px;
            right: 30px;
            background: linear-gradient(135deg, #ff416c, #ff4b2b);
            color: white;
            padding: 15px 25px;
            border-radius: 10px;
            box-shadow: 0 10px 25px rgba(0, 0, 0, 0.2);
            display: flex;
            align-items: center;
            gap: 15px;
            transform: translateX(150%);
            transition: transform 0.5s cubic-bezier(0.68, -0.55, 0.265, 1.55);
            z-index: 1000;
        }

        .error-message.show {
            transform: translateX(0);
        }

        .message-icon {
            font-size: 24px;
        }

        @keyframes float {
            0%, 100% {
                transform: translateY(0);
            }
            50% {
                transform: translateY(-15px);
            }
        }

        @keyframes float-element {
            0% {
                transform: translateY(0) rotate(0deg);
                opacity: 0.7;
            }
            33% {
                transform: translateY(-30px) rotate(120deg);
                opacity: 0.4;
            }
            66% {
                transform: translateY(20px) rotate(240deg);
                opacity: 0.8;
            }
            100% {
                transform: translateY(0) rotate(360deg);
                opacity: 0.7;
            }
        }

        @keyframes shimmer {
            0% {
                background-position: 200% 0;
            }
            100% {
                background-position: -200% 0;
            }
        }

        @keyframes fadeIn {
            from {
                opacity: 0;
                transform: translateY(10px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }

        @keyframes shake {
            0%, 100% { transform: translateX(0); }
            10%, 30%, 50%, 70%, 90% { transform: translateX(-5px); }
            20%, 40%, 60%, 80% { transform: translateX(5px); }
        }

        .shake {
            animation: shake 0.5s ease-in-out;
        }

        @keyframes card-entry {
            0% {
                opacity: 0;
                transform: translateY(30px) rotateX(30deg);
            }
            100% {
                opacity: 1;
                transform: translateY(0) rotateX(0);
            }
        }

        .upload-card {
            animation: card-entry 0.8s cubic-bezier(0.175, 0.885, 0.32, 1.275) forwards;
        }

        @media (max-width: 600px) {
            .upload-card {
                padding: 30px 25px;
            }

            .header h1 {
                font-size: 24px;
            }

            .upload-text {
                font-size: 16px;
            }

            .upload-area {
                padding: 30px 15px;
            }
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="upload-card">
            <div class="floating-elements">
                <div class="element"></div>
                <div class="element"></div>
                <div class="element"></div>
                <div class="element"></div>
            </div>

            <div class="header">
                <div class="logo">
                    <i class="fas fa-cloud-upload-alt"></i>
                </div>
                <h1>آپلود فایل</h1>
                <p>فایل خود را انتخاب کنید یا آن را در این ناحیه رها کنید. حداکثر حجم فایل: 10 مگابایت</p>
            </div>

            <form enctype="multipart/form-data" action="/upload" method="POST" id="uploadForm">
                <div class="upload-area" id="uploadArea">
                    <div class="upload-icon">
                        <i class="fas fa-cloud-upload-alt"></i>
                    </div>
                    <div class="upload-text">فایل خود را اینجا رها کنید یا برای انتخاب کلیک کنید</div>
                    <div class="upload-subtext">پشتیبانی از فایل‌های: تصاویر (JPG, PNG, GIF)، اسناد (PDF, DOC)، و سایر فایل‌ها</div>
                    <button type="button" class="browse-btn" id="browseBtn">
                        <i class="fas fa-folder-open"></i> انتخاب فایل
                    </button>
                    <input type="file" name="myFile" class="file-input" id="fileInput" accept=".jpg,.jpeg,.png,.gif,.pdf,.doc,.docx,.txt">
                </div>

                <div class="file-info" id="fileInfo">
                    <div class="file-preview" id="filePreview">
                        <!-- Preview will be inserted here -->
                    </div>

                    <div class="file-details">
                        <div class="file-detail">
                            <span class="detail-label">نام فایل:</span>
                            <span class="detail-value" id="fileName">-</span>
                        </div>
                        <div class="file-detail">
                            <span class="detail-label">نوع فایل:</span>
                            <span class="detail-value" id="fileType">-</span>
                        </div>
                        <div class="file-detail">
                            <span class="detail-label">حجم فایل:</span>
                            <span class="detail-value" id="fileSize">-</span>
                        </div>
                        <div class="file-detail">
                            <span class="detail-label">تاریخ آخرین تغییر:</span>
                            <span class="detail-value" id="fileDate">-</span>
                        </div>
                    </div>
                </div>

                <div class="progress-bar" id="progressBar">
                    <div class="progress-fill" id="progressFill"></div>
                </div>

                <button type="submit" class="submit-btn" id="submitBtn" disabled>
                    <i class="fas fa-upload"></i> آپلود فایل
                </button>
            </form>
        </div>
    </div>

    <div class="success-message" id="successMessage">
        <i class="fas fa-check-circle message-icon"></i>
        <div>فایل با موفقیت آپلود شد!</div>
    </div>

    <div class="error-message" id="errorMessage">
        <i class="fas fa-exclamation-circle message-icon"></i>
        <div id="errorText">خطایی در آپلود فایل رخ داد</div>
    </div>

    <script>
        document.addEventListener('DOMContentLoaded', function() {
            const uploadArea = document.getElementById('uploadArea');
            const fileInput = document.getElementById('fileInput');
            const browseBtn = document.getElementById('browseBtn');
            const fileInfo = document.getElementById('fileInfo');
            const filePreview = document.getElementById('filePreview');
            const fileName = document.getElementById('fileName');
            const fileType = document.getElementById('fileType');
            const fileSize = document.getElementById('fileSize');
            const fileDate = document.getElementById('fileDate');
            const submitBtn = document.getElementById('submitBtn');
            const progressBar = document.getElementById('progressBar');
            const progressFill = document.getElementById('progressFill');
            const successMessage = document.getElementById('successMessage');
            const errorMessage = document.getElementById('errorMessage');
            const errorText = document.getElementById('errorText');
            const uploadCard = document.querySelector('.upload-card');
            const uploadForm = document.getElementById('uploadForm');

            // انیمیشن ورود اولیه
            setTimeout(() => {
                uploadCard.style.opacity = '1';
            }, 100);

            // رویداد کلیک روی ناحیه آپلود
            browseBtn.addEventListener('click', function() {
                fileInput.click();
            });

            uploadArea.addEventListener('click', function(e) {
                if (e.target !== browseBtn) {
                    fileInput.click();
                }
            });

            // رویداد تغییر فایل
            fileInput.addEventListener('change', function() {
                if (this.files && this.files[0]) {
                    handleFileSelection(this.files[0]);
                }
            });

            // Drag and drop functionality
            ['dragenter', 'dragover', 'dragleave', 'drop'].forEach(eventName => {
                uploadArea.addEventListener(eventName, preventDefaults, false);
            });

            function preventDefaults(e) {
                e.preventDefault();
                e.stopPropagation();
            }

            ['dragenter', 'dragover'].forEach(eventName => {
                uploadArea.addEventListener(eventName, highlight, false);
            });

            ['dragleave', 'drop'].forEach(eventName => {
                uploadArea.addEventListener(eventName, unhighlight, false);
            });

            function highlight() {
                uploadArea.classList.add('dragover');
            }

            function unhighlight() {
                uploadArea.classList.remove('dragover');
            }

            uploadArea.addEventListener('drop', handleDrop, false);

            function handleDrop(e) {
                const dt = e.dataTransfer;
                const files = dt.files;

                if (files.length > 0) {
                    handleFileSelection(files[0]);
                }
            }

            // مدیریت انتخاب فایل
            function handleFileSelection(file) {
                // اعتبارسنجی حجم فایل (حداکثر 10 مگابایت)
                const maxSize = 10 * 1024 * 1024; // 10MB

                if (file.size > maxSize) {
                    showError('حجم فایل باید کمتر از 10 مگابایت باشد');
                    uploadCard.classList.add('shake');
                    setTimeout(() => {
                        uploadCard.classList.remove('shake');
                    }, 500);
                    return;
                }

                // نمایش اطلاعات فایل
                displayFileInfo(file);

                // فعال کردن دکمه ارسال
                submitBtn.disabled = false;
            }

            // نمایش اطلاعات فایل
            function displayFileInfo(file) {
                // نمایش بخش اطلاعات
                fileInfo.classList.add('show');

                // تنظیم اطلاعات فایل
                fileName.textContent = file.name;
                fileType.textContent = getFileType(file.type);
                fileSize.textContent = formatFileSize(file.size);
                fileDate.textContent = formatDate(new Date(file.lastModified));

                // نمایش پیش‌نمایش
                displayPreview(file);
            }

            // نمایش پیش‌نمایش
            function displayPreview(file) {
                filePreview.innerHTML = '';

                if (file.type.startsWith('image/')) {
                    const reader = new FileReader();
                    reader.onload = function(e) {
                        const img = document.createElement('img');
                        img.src = e.target.result;
                        img.className = 'preview-image';
                        img.alt = 'پیش‌نمایش تصویر';
                        filePreview.appendChild(img);
                    };
                    reader.readAsDataURL(file);
                } else {
                    // آیکون بر اساس نوع فایل
                    const icon = document.createElement('i');
                    icon.className = 'preview-icon';

                    if (file.type.includes('pdf')) {
                        icon.className += ' fas fa-file-pdf';
                    } else if (file.type.includes('document') || file.type.includes('msword') || file.name.endsWith('.doc') || file.name.endsWith('.docx')) {
                        icon.className += ' fas fa-file-word';
                    } else if (file.type.includes('text')) {
                        icon.className += ' fas fa-file-alt';
                    } else {
                        icon.className += ' fas fa-file';
                    }

                    filePreview.appendChild(icon);

                    const fileTypeText = document.createElement('div');
                    fileTypeText.textContent = getFileType(file.type);
                    fileTypeText.style.marginTop = '10px';
                    fileTypeText.style.color = '#666';
                    filePreview.appendChild(fileTypeText);
                }
            }

            // دریافت نوع فایل به صورت فارسی
            function getFileType(mimeType) {
                if (mimeType.startsWith('image/')) {
                    return 'تصویر';
                } else if (mimeType.includes('pdf')) {
                    return 'سند PDF';
                } else if (mimeType.includes('document') || mimeType.includes('msword')) {
                    return 'سند ورد';
                } else if (mimeType.includes('text')) {
                    return 'فایل متنی';
                } else {
                    return mimeType || 'ناشناخته';
                }
            }

            // فرمت حجم فایل
            function formatFileSize(bytes) {
                if (bytes === 0) return '0 بایت';

                const k = 1024;
                const sizes = ['بایت', 'کیلوبایت', 'مگابایت', 'گیگابایت'];
                const i = Math.floor(Math.log(bytes) / Math.log(k));

                return parseFloat((bytes / Math.pow(k, i)).toFixed(2)) + ' ' + sizes[i];
            }

            // فرمت تاریخ
            function formatDate(date) {
                const options = {
                    year: 'numeric',
                    month: 'long',
                    day: 'numeric',
                    hour: '2-digit',
                    minute: '2-digit'
                };
                return new Intl.DateTimeFormat('fa-IR', options).format(date);
            }

            // نمایش پیام خطا
            function showError(message) {
                errorText.textContent = message;
                errorMessage.classList.add('show');

                setTimeout(() => {
                    errorMessage.classList.remove('show');
                }, 5000);
            }

            // نمایش پیام موفقیت
            function showSuccess(message) {
                successMessage.querySelector('div').textContent = message || 'فایل با موفقیت آپلود شد!';
                successMessage.classList.add('show');

                setTimeout(() => {
                    successMessage.classList.remove('show');
                }, 5000);
            }

            // شبیه‌سازی آپلود
            uploadForm.addEventListener('submit', function(e) {
                e.preventDefault();

                if (!fileInput.files[0]) {
                    showError('لطفاً یک فایل انتخاب کنید');
                    return;
                }

                // نمایش نوار پیشرفت
                progressBar.style.display = 'block';
                submitBtn.disabled = true;

                // شبیه‌سازی آپلود با انیمیشن پیشرفت
                let progress = 0;
                const interval = setInterval(() => {
                    progress += 5;
                    progressFill.style.width = progress + '%';

                    if (progress >= 100) {
                        clearInterval(interval);

                        // نمایش پیام موفقیت
                        showSuccess();

                        // مخفی کردن نوار پیشرفت پس از 2 ثانیه
                        setTimeout(() => {
                            progressBar.style.display = 'none';
                            progressFill.style.width = '0%';

                            // بازنشانی فرم
                            uploadForm.reset();
                            fileInfo.classList.remove('show');
                            submitBtn.disabled = true;
                        }, 2000);
                    }
                }, 100);
            });
        });
    </script>
</body>
</html>
		`
		fmt.Fprint(w, html)
		return
	}
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Println("Error retriveing the file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	fmt.Printf("File Size: %+v\n", handler.Size)
	fmt.Printf("MIME Header: %+v\n", handler.Header)

	dst, err := os.Create("./temp-images/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Successfully uploaded File: %s", handler.Filename)
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	fmt.Println("Server is running on http://localhost:8080/upload")
	http.ListenAndServe(":8080", nil)
}
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func profilePictureHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		html := `
			<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Upload Profile Picture</title>
		</head>
		<body>
			<h2>Upload Profile Picture</h2>
			<form action="/profile-picture" method="POST" enctype="multipart/form-data">
				<input type="file" name="picture" accept="image/*">
				<br><br>
				<button type="submit">Upload</button>
			</form>
		</body>
		</html>
		`

		fmt.Fprint(w, html)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseMultipartForm(5 << 20)
		if err != nil {
			http.Error(w, "File is too large", http.StatusBadRequest)
			return
		}

		file, handler, err := r.FormFile("picture")
		if err != nil {
			http.Error(w, "Error reading file", http.StatusBadRequest)
			return
		}
		defer file.Close()

		if handler.Size > 5<<20 {
			fmt.Fprint(w, "File is too large")
			return
		}

		err = os.MkdirAll("upload", os.ModePerm)
		if err != nil {
			http.Error(w, "Cannot create upload directory", http.StatusInternalServerError)
			return
		}
		dst, err := os.Create("uploads/" + handler.Filename)
		if err != nil {
			http.Error(w, "Cannot save file", http.StatusInternalServerError)
			return
		}

		defer dst.Close()
		_, err = io.Copy(dst, file)
		if err != nil {
			http.Error(w, "Error saving file", http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/profile-picture", profilePictureHandler)

	fmt.Println("Server running on http://localhost:8080/profile-picture")
	http.ListenAndServe(":8080", nil)
}
