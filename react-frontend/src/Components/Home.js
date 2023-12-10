import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { EditOutlined, SettingOutlined } from '@ant-design/icons';
import { Avatar, Card, Empty } from 'antd';
import data from '../logos/data.json';
import NavBar from './NavBar';
import '../styles/home.css';
const { Meta } = Card;

function Home() {
  const navigate = useNavigate();
  const [enterClass, setEnterClass] = useState('');

  // Add a useEffect to trigger the initial animation
  useEffect(() => {
    setEnterClass('car-card-enter');
  }, []);

  const [cars, setCars] = useState([]);
  const [carImage, setCarImage] = useState('');
  const token = localStorage.getItem('token');
  const username = localStorage.getItem('username');
  const pageId = 1;
  const pageSize = 10;
  const url = `http://localhost:8080/cars/users/${username}?page_id=${pageId}&page_size=${pageSize}`;

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch(url, {
          method: 'GET',
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });

        if (response.ok) {
          const data = await response.json();
          setCars(data);
        } else {
          console.error('Error fetching data:', response.statusText);
        }
      } catch (error) {
        console.error('Error during data fetch:', error);
      }
    };

    fetchData();
  }, [token, url]);

  useEffect(() => {
    // Check if cars is not empty before accessing it
    if (cars.length > 0) {
      const carResult = data.filter((item) =>
        item.name.toLowerCase().includes(carName.toLowerCase())
      );

      // Check if carResult is not empty before accessing it
      if (carResult.length > 0) {
        setCarImage(carResult[0].image.thumb);
      }
    }
  }, [cars]);

  const carName = 'ferrari'; // Move carName here

  const handleCreateCarClick = () => {
    const url = `/home/${username}/create-car`;
    navigate(url);
  };

  const handleToLogMaintenancePath = (car) => {
    const url = `/home/${username}/${car.vin}/log-maintenance`;
    navigate(url, { state: { car } });
  };

  return (
    <div className='home'>
      <NavBar />
      <div className='create-button-div'>
        <button className='create-car-button' onClick={handleCreateCarClick}>
          Add a Car
        </button>
      </div>

      {cars.length === 0 ? (
        <Empty image={Empty.PRESENTED_IMAGE_SIMPLE} />
      ) : (
        <div className='home-page-title-cars-container'>
          <div className='cars-containers'>
            <h1 className='homepage-title'>{username}'s' cars</h1>
            {cars.map((car) => (
              <Card
                className='car-card'
                key={car.vin}
                style={{
                  width: 300,
                }}
                actions={[
                  <div>
                    <button
                      className='log-maintenance-button'
                      onClick={() => handleToLogMaintenancePath(car)}
                    >
                      Delete Car
                    </button>
                  </div>,
                  <div className='card-button'>
                    <button
                      className='log-maintenance-button'
                      onClick={() => handleToLogMaintenancePath(car)}
                    >
                      Log Maintenance
                    </button>
                  </div>,
                  <div>
                    <button
                      className='log-maintenance-button'
                      onClick={() => handleToLogMaintenancePath(car)}
                    >
                      Show Timeline
                    </button>
                  </div>,
                ]}
              >
                <Meta
                  className='hello'
                  avatar={<Avatar src={carImage} />}
                  title={'VIN: ' + car.vin}
                  description={car.make + ' ' + car.model + ' ' + car.year}
                />
              </Card>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}

export default Home;
