import React from 'react';
const Navbar: React.FC = () => {
  return (
    <nav className='navbar'>
      <ul>
        <li><a href="/">Home</a></li>
        <li><a href="/userspage">Users</a></li>
      </ul>
    </nav>
  );
};

export default Navbar;