import { NavLink, Link } from "react-router-dom";

const Header = () => {
  return (
    <div className="header">
      <Link to="" className="logo">
        AreTheyUp?
      </Link>
      <div className="header-right">
        <NavLink to="">Home</NavLink>
        <NavLink to="/about">About</NavLink>
        <NavLink to="/contact">Contact</NavLink>
        <NavLink to="/reportissue">REPORT ISSUE</NavLink>
      </div>
    </div>
  );
};

export default Header;
