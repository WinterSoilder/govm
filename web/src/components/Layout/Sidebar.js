/* eslint-disable jsx-a11y/anchor-is-valid */
import "./Layout.css";
import logo from "../../assets/logo.png";
import { Link } from "react-router-dom";

const Sidebar = ({ sidebarOpen, closeSidebar }) => {
  return (
    <div className={sidebarOpen ? "sidebar_responsive" : ""} id="sidebar">
      <div className="sidebar__title">
        <div className="sidebar__img">
          <img src={logo} alt="logo" />
          <h1 style={{ fontSize: '34px' }}>GO VM</h1>
        </div>
        <i
          onClick={() => closeSidebar()}
          className="fa fa-times"
          id="sidebarIcon"
          aria-hidden="true"
        ></i>
      </div>

      <div className="sidebar__menu">
        <div className="sidebar__link active_menu_link">
          <Link
            to={'/'}
          >
            Dashboard
          </Link>
        </div>
        <div className="sidebar__link">
          <Link
            to={'/vms'}
          >
            <i className="fa fa-user-secret" aria-hidden="true"></i>
            <div>
              My Virtual Machines
            </div>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default Sidebar;