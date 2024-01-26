import { Button, Form, Input, message } from 'antd';
import React from 'react';
import { useNavigate } from 'react-router-dom';
import '../styles/Login.css'; // Import the external styles
import NavBar from './NavBar';

const Login = () => {
  const navigate = useNavigate();

  const handleBackToLandingPageClick = () => {
    navigate('/');
  };

  const handleLoginFailed = () => {
    // console.log('LOGIN FAILED');
  };

  const handleLogin = async (values) => {
    // 'values' contains the form field values (including username and password)

    try {
      const response = await fetch('http://localhost:8080/users/login', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(values),
      });

      if (response.ok) {
        const data = await response.json();
        const token = data['access_token'];
        const username = data['user']['username'];
        localStorage.setItem('token', token);
        localStorage.setItem('username', username);

        const successMessage = `Successfully logged in as ${data['user']['username']}`;
        message.success(successMessage);
        // Redirect to the home page on successful login
        navigate(`/home/${username}`);
      } else {
        console.error('Authentication failed');
        const errorMessage = `Error, Password or Username is incorrect`;
        message.error(errorMessage);
      }
    } catch (error) {
      console.error('Error during login', error);
    }
  };

  return (
    <div className='login-page'>
      <NavBar />
      <div className='login-container'>
        <div className='login-titles'>
          <h1>Welcome,</h1>
          <h2>Please Log in</h2>
        </div>
        <div className='form-container'>
          <Form
            className='login-form'
            name='basic'
            labelCol={{
              span: 8,
            }}
            wrapperCol={{
              span: 16,
            }}
            initialValues={{
              remember: true,
            }}
            onFinish={handleLogin}
            onFinishFailed={handleLoginFailed}
            autoComplete='off'
          >
            <Form.Item
              className='login-form-item'
              label='Username'
              name='username'
              rules={[
                {
                  required: true,
                  message: 'Please input your username!',
                },
              ]}
            >
              <Input />
            </Form.Item>

            <Form.Item
              className='login-form-item'
              label='Password'
              name='password'
              rules={[
                {
                  required: true,
                  message: 'Please input your password!',
                },
              ]}
            >
              <Input.Password />
            </Form.Item>

            <Form.Item
              wrapperCol={{
                offset: 8,
                span: 16,
              }}
            >
              <div className='login-form-buttons'>
                <Button
                  className='login-button'
                  type='primary'
                  htmlType='submit'
                >
                  Sign in
                </Button>
                <button
                  className='back-to-home'
                  onClick={handleBackToLandingPageClick}
                >
                  EXIT
                </button>
              </div>
            </Form.Item>
          </Form>
        </div>
      </div>
    </div>
  );
};

export default Login;
