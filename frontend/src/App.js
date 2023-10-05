import "./App.css";
import Header from "./components/shared/Header";
import About from "./components/shared/About";
import Contact from "./components/shared/Contact";
import ReportIssue from "./components/shared/ReportIssue";
import Content from "./components/Content";
import { Routes, Route, Outlet } from "react-router-dom";

function App() {
  return (
    <div className="App">
      <Header />
      <Outlet />
      <Routes>
        <Route path="" element={<Content />}></Route>
        <Route path="/about" element={<About />}></Route>
        <Route path="/contact" element={<Contact />}></Route>
        <Route path="/reportissue" element={<ReportIssue />}></Route>
      </Routes>
    </div>
  );
}

export default App;
