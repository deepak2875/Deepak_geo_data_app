import { Link } from 'react-router-dom';

function Navbar() {
  return (
    <nav>
      <Link to="/">Login</Link>
      <Link to="/register">Register</Link>
      <Link to="/map">Map</Link>
      <Link to="/upload">Upload File</Link>
    </nav>
  );
}

export default Navbar;


