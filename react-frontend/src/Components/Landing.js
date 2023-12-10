import { useNavigate } from 'react-router-dom';
import React, { useEffect, useState } from 'react';
import NavBar from './NavBar';

function Landing() {
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const originalTitle = 'Car Records in one place'; // Your original title

  useEffect(() => {
    let index = 0;
    const titleInterval = setInterval(() => {
      setTitle(originalTitle.substring(0, index));
      index++;
      if (index > originalTitle.length) {
        clearInterval(titleInterval);
      }
    }, 100); // Adjust the interval time as needed
    return () => clearInterval(titleInterval); // Cleanup on component unmount
  }, []);

  const handleSignInClick = () => {
    navigate('/users/login');
  };

  const handleCreateAccountClick = () => {
    navigate('/users/create');
  };

  return (
    <div className='landing'>
      <NavBar />
      <div className='landing-container'>
        <h1 className='title'>{title}</h1>
        <div className='login-buttons'>
          <button className='button-3' onClick={handleSignInClick}>
            Sign In
          </button>
          <button className='button-12' onClick={handleCreateAccountClick}>
            Create Account
          </button>
        </div>
      </div>
    </div>
  );
}

export default Landing;
