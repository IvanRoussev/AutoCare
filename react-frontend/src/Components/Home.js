import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';
import { Avatar, Card, message } from 'antd';
import data from '../logos/data.json';
import NavBar from './NavBar';
import '../styles/home.css';
const { Meta } = Card;

function Home() {
  const location = useLocation();
  const { image } = location.state || {};
  const navigate = useNavigate();

  const [cars, setCars] = useState([]);
  const [carImage, setCarImage] = useState('');

  const token = localStorage.getItem('token');
  const username = localStorage.getItem('username');
  const pageId = 1;
  const pageSize = 10;

  useEffect(() => {
    const fetchData = async () => {
      const url = `http://localhost:8080/cars/users/${username}?page_id=${pageId}&page_size=${pageSize}`;
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
  }, [token, username]);

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

  const carName = 'Ferrari'; // Move carName here

  const handleCreateCarClick = () => {
    const url = `/home/${username}/create-car`;
    navigate(url);
  };

  const handleToLogMaintenancePath = (car) => {
    const url = `/home/${username}/${car.vin}/log-maintenance`;
    navigate(url, { state: { car } });
  };

  const navigateToTimeline = (car) => {
    const url = `/home/${username}/${car.vin}/timeline`;
    navigate(url, { state: { car } });
  };

  const deleteCar = async (car) => {
    const url = `http://localhost:8080/cars/vin/${car.vin}`;
    try {
      const response = await fetch(url, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (response.ok) {
        // Remove the deleted car from the state
        setCars((prevCars) => prevCars.filter((c) => c.vin !== car.vin));
        const successMessage = `Successfully deleted ${car.year} ${car.make} ${car.model}`;
        message.success(successMessage);
      } else {
        console.error('Error deleting car:', response.statusText);
        const errorMessage = `Error could not delete the car ${car.year} ${car.make} ${car.model}`;
        message.error(errorMessage);
      }
    } catch (error) {
      console.error('Error during car deletion:', error);
      const errorMessage = `Error could not delete the car ${car.year} ${car.make} ${car.model}`;
      message.error(errorMessage);
    }
  };

  return (
    <div className='home'>
      <NavBar />
      <div className='username-title-homepage'>
        <h1 className='homepage-title'>{username}'s' cars</h1>
      </div>
      <div className='create-button-div'>
        <button className='create-car-button' onClick={handleCreateCarClick}>
          Add a Car
        </button>
      </div>
      <div className='home-page-title-cars-container'>
        <div className='cars-containers'>
          {cars.map((car) => (
            <Card
              className='car-card'
              key={car.vin}
              style={{
                width: 300,
              }}
              actions={[
                <div className='card-button'>
                  <button
                    className='delete-car-button'
                    onClick={() => deleteCar(car)}
                  >
                    Delete Car
                  </button>
                </div>,
                <div>
                  <button
                    className='log-maintenance-button'
                    onClick={() => handleToLogMaintenancePath(car)}
                  >
                    Log Maintenance
                  </button>
                </div>,
                <div>
                  <button
                    className='create-timeline-button'
                    onClick={() => navigateToTimeline(car)}
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
    </div>
  );
}

export default Home;
