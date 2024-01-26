// App.jsx

import React from 'react';
import Landing from './Components/Landing';
import './App.css';
import Login from './Components/Login';
import CreateUser from './Components/CreateUser';
import Home from './Components/Home';
import CreateCar from './Components/Car/GetMake';
import GetCarInfo from './Components/Car/GetCarInfo';
import LogMaintenance from './Components/Maintenance/LogMaintenance';
import ShowTimeline from './Components/Maintenance/ShowTimeline';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {
  return (
    <Router>
      <div>
        <Routes>
          <Route path='/' element={<Landing />} />
          <Route path='/users/login' element={<Login />} />
          <Route path='/users/create' element={<CreateUser />} />
          <Route path='/home/:username' element={<Home />} />
          <Route path='/home/:username/create-car' element={<CreateCar />} />
          <Route
            path='/home/:username/create-car/:name'
            element={<GetCarInfo />}
          />
          <Route
            path='home/:username/:car_vin/log-maintenance'
            element={<LogMaintenance />}
          />
          <Route
            path='home/:username/:car_vin/timeline'
            element={<ShowTimeline />}
          />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
