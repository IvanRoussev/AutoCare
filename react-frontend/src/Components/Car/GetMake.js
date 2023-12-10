import { Card, Col, Row, Input } from 'antd';
import { Link } from 'react-router-dom';
import { useEffect, useState } from 'react';
import data from '../../logos/data.json';
const { Search } = Input;

function CreateCar() {
  const [searchQuery, setSearchQuery] = useState('');
  const [filteredData, setFilteredData] = useState([]);
  const username = localStorage.getItem('username');

  useEffect(() => {
    const filteredResults = data.filter((item) =>
      item.name.toLowerCase().includes(searchQuery.toLocaleLowerCase())
    );

    setFilteredData(filteredResults);
  }, [searchQuery]);

  const handleSearch = (e) => {
    setSearchQuery(e.target.value);
  };

  return (
    <div>
      <Search
        placeholder='Search car manufacturer'
        value={searchQuery}
        onChange={handleSearch}
      />
      <Row gutter={16}>
        {filteredData.map((item) => (
          <Col key={item.name} span={8}>
            <div className='card-logo-cards'>
              <Link
                to={`/home/${username}/create-car/${item.name}`}
                make={item.name}
              >
                <Card title={item.name} bordered={false}>
                  <img
                    src={item.image.optimized}
                    alt={item.name}
                    style={{ width: '100%' }}
                  />
                </Card>
              </Link>
            </div>
          </Col>
        ))}
      </Row>
    </div>
  );
}

export default CreateCar;
