import React, { createContext, useContext, useMemo } from 'react';
import { Form, Input, Button, message } from 'antd';
import { useNavigate } from 'react-router-dom';
import '../styles/CreateUser.css';

const MyFormItemContext = createContext([]);

function toArr(str) {
  return Array.isArray(str) ? str : [str];
}

const MyFormItemGroup = ({ prefix, children }) => {
  const prefixPath  = useContext(MyFormItemContext);
  const concatPath = useMemo(
    () => [...prefixPath, ...toArr(prefix)],
    [prefixPath, prefix]
  );

  return (
    <MyFormItemContext.Provider value={concatPath}>
      {children}
    </MyFormItemContext.Provider>
  );
};

const MyFormItem = ({ name, ...props }) => {
  const prefixPath = useContext(MyFormItemContext);
  const concatName =
    name !== undefined ? [...prefixPath, ...toArr(name)] : undefined;

  return <Form.Item name={concatName} {...props} />;
};
const CreateUser = () => {
  const navigate = useNavigate();

  const handleBackToLandingPageClick = () => {
    navigate('/');
  };

  const onFinish = async (values) => {
    const user = values['user'];
    try {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        body: JSON.stringify({
          full_name: user['fullname'],
          username: user['username'],
          password: user['password'],
          email: user['email'],
        }),
        headers: {
          'Content-type': 'application/json',
        },
      });
      if (response.ok) {
        const successMessage = `Successfully Created User ${user['username']}`;
        message.success(successMessage);

        navigate('/users/login');

        // Optionally, you can redirect or perform other actions upon successful creation
      } else {
        console.error('Failed to create user:', response.statusText);
      }
    } catch (error) {
      console.error('Error during POST request:', error);
    }
  };

  return (
    <div className='create-user-container'>
      <Form
        className='create-user-form'
        name='form_item_path'
        layout='vertical'
        onFinish={onFinish}
      >
        <MyFormItemGroup prefix={['user']}>
          <MyFormItem
            name='fullname'
            label='Full Name'
            rules={[
              {
                required: true,
                message: 'Please enter your Full Name',
              },
            ]}
          >
            <Input />
          </MyFormItem>
          <MyFormItem
            name='username'
            label='Username'
            rules={[
              {
                required: true,
                message: 'Please enter a username',
              },
            ]}
          >
            <Input />
          </MyFormItem>
          <MyFormItem
            name='password'
            label='Password'
            rules={[
              {
                required: true,
                message: 'Please enter your a Password',
              },
            ]}
          >
            <Input.Password />
          </MyFormItem>
          <MyFormItem
            name='email'
            label='Email'
            rules={[
              {
                type: 'email',
                message: 'Please enter a valid email address',
              },
              {
                required: true,
                message: 'Please enter your email address',
              },
            ]}
          >
            <Input />
          </MyFormItem>
        </MyFormItemGroup>
        <div className='create-user-buttons-div'>
          <button
            className='back-to-home-from-createuser'
            onClick={handleBackToLandingPageClick}
          >
            Back to Home
          </button>
          <Button type='primary' htmlType='submit'>
            Submit
          </Button>
        </div>
      </Form>
    </div>
  );
};

export default CreateUser;
