import React, { useState, useEffect } from 'react';
import { Button, Form, Input, Space, message } from 'antd';
import { useParams, useNavigate } from 'react-router-dom';

const token = localStorage.getItem('token');
const username = localStorage.getItem('username');

const CreateCarAPI = async (car) => {
  const vin = car['vin'];
  const make = car['make'];
  const model = car['model'];
  const year = parseInt(car['year']);
  try {
    const response = await fetch('http://localhost:8080/cars', {
      method: 'POST',
      body: JSON.stringify({
        vin: vin,
        make: make,
        model: model,
        year: year,
      }),
      headers: {
        'Content-type': 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });
    if (response.ok) {
      const successMessage = `Successfully Added Your ${make}  ${model}`;
      message.success(successMessage);
    } else {
      console.error('Failed to create car:', response.statusText);
    }
  } catch (error) {
    console.error('Error during POST request:', error);
  }
};

function SubmitButton({ form, onFinish }) {
  const [submittable, setSubmittable] = useState(false);
  const values = Form.useWatch([], form);

  useEffect(() => {
    form
      .validateFields({
        validateOnly: true,
      })
      .then(
        () => {
          setSubmittable(true);
        },
        () => {
          setSubmittable(false);
        }
      );
  }, [values, form]);

  return (
    <Button
      type='primary'
      htmlType='submit'
      disabled={!submittable}
      onClick={onFinish}
    >
      Submit
    </Button>
  );
}

function GetCarInfo() {
  const navigate = useNavigate();
  const url = `/home/${username}`;

  const { name } = useParams();

  var [form] = Form.useForm();
  const onFinish = (values) => {
    values.make = name;

    // Simulate API submission success
    CreateCarAPI(values);
    navigate(url);
  };

  return (
    <Form
      form={form}
      name='validateOnly'
      layout='vertical'
      autoComplete='off'
      onFinish={onFinish}
    >
      <Form.Item
        name='vin'
        label='VIN'
        rules={[
          {
            required: true,
            message: 'Please enter VIN',
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name='model'
        label='Model'
        rules={[
          {
            required: true,
            message: 'Please enter the Model of car',
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item
        name='year'
        label='Year'
        rules={[
          {
            required: true,
            message: 'Please enter the Year of your car',
          },
        ]}
      >
        <Input />
      </Form.Item>
      <Form.Item>
        <Space>
          <SubmitButton form={form} />
          <Button htmlType='reset'>Reset</Button>
        </Space>
      </Form.Item>
    </Form>
  );
}

export default GetCarInfo;
