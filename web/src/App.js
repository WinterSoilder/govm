import { useState } from "react";
import Main from "./components/Main/Main";
import Navbar from "./components/Layout/Navbar";
import Sidebar from "./components/Layout/Sidebar";
import ALLVms from "./components/VMConfig/AllVMs";
import VMFormIndex from "./components/VMConfig/FormIndex";
import { Route, Link, Routes } from "react-router-dom";
import Login from "./components/Auth/Login";

const App = () => {
  const [sidebarOpen, setsidebarOpen] = useState(false);
  const openSidebar = () => {
    setsidebarOpen(true);
  };
  const closeSidebar = () => {
    setsidebarOpen(false);
  };
  return (
    <div className="container">
      
      <div className="App">
        <Routes>
          <Route path="/" element={<Navbar sidebarOpen={sidebarOpen} openSidebar={openSidebar} />}>
            <Route index element={<Main />} />
            <Route path="login" element={<Login />} />
            <Route path="vms" element={<ALLVms />} />
            <Route path="vms/add" element={<VMFormIndex />} />
          </Route>
        </Routes>
      </div>
      <Sidebar sidebarOpen={sidebarOpen} closeSidebar={closeSidebar} />
    </div>
  );
};

export default App;