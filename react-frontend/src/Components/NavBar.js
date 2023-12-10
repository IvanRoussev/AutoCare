function NavBar() {
  return (
    <nav>
      <div className='nav-wrapper'>
        <a href='#' className='brand-logo'>
          Auto Records Hub
        </a>
        <ul id='menu' className='menu'>
          <li>
            <a href=''>Home</a>
          </li>
          <li>
            <a href=''>About</a>
          </li>
          <li>
            <a href=''>Github</a>
          </li>
        </ul>
      </div>
    </nav>
  );
}

export default NavBar;
