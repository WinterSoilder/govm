/* eslint-disable jsx-a11y/anchor-is-valid */
import "./Layout.css";
import avatar from "../../assets/avatar.svg";
import { Outlet } from "react-router-dom";

const Navbar = ({ sidebarOpen, openSidebar, closeSidebar }) => {
  return (
    <div>
    <nav className="navbar">
      <div className="nav_icon" onClick={() => openSidebar()}>
        <i className="fa fa-bars" aria-hidden="true"></i>
      </div>
      <div className="navbar__left">
       
        <a className="active_link" href="#">
          Admin
        </a>
      </div>
      <div className="navbar__right">
        <a href="#!">
          <img width="30" src={avatar} alt="avatar" />
        </a>
      </div>
    </nav>
    <Outlet />
    </div>
  );
};

export default Navbar;