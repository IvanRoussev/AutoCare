import { Card, Col, Row, Input } from 'antd';
import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import data from '../../logos/data.json';
import '../../styles/getMake.css';
const { Search } = Input;

function CreateCar() {
  const navigate = useNavigate();
  const [searchQuery, setSearchQuery] = useState('');
  const [filteredData, setFilteredData] = useState([]);
  const username = localStorage.getItem('username');

  useEffect(() => {
    const filteredResults = data.filter((item) =>
      item.name.tocLowerCase().includes(searchQuery.toLocaleLowerCase())
    );

    setFilteredData(filteredResults);
  }, [searchQuery]);

  const handleSearch = (e) => {
    setSearchQuery(e.target.value);
  };

  const handleCreateCarClick = (item) => {
    let image = item.image.thumb;
    const url = `/home/${username}/create-car/${item.name}`;
    navigate(url, { state: { image } });
  };

  const handleBacktoHomeScreen = () => {
    let url = `/home/${username}`;
    navigate(url);
  };

  return (
    <div className='create-car-page'>
      <button className='esc-button' onClick={handleBacktoHomeScreen}>
        ESC
      </button>
      <div className='search-container'>
        <h1 className='getMake-title'>Search for your cars manufacturer</h1>
        <Search
          className='search-bar'
          placeholder='Search car manufacturer'
          value={searchQuery}
          onChange={handleSearch}
        />
      </div>
      <Row gutter={16}>
        <div className='card-container'>
          {filteredData.map((item) => (
            <Col key={item.name} xs={8} sm={6} md={5} lg={4} xl={3}>
              <div className='card-logo-cards'>
                <Link to={`/home/${username}/create-car/${item.name}`}>
                  <Card
                    title={item.name}
                    bordered={false}
                    className='card-make'
                    onClick={() => handleCreateCarClick(item)}
                  >
                    <img
                      className='make-image'
                      src={item.image.optimized}
                      alt={item.name}
                    />
                  </Card>
                </Link>
              </div>
            </Col>
          ))}
        </div>
      </Row>
    </div>
  );
}

export default CreateCar;
