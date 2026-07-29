import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "../pages/Home/Home";
import Playground from "../pages/Playground/Playground";

export default function AppRouter() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route
          path="/playground"
          element={<Playground />}
        />
      </Routes>
    </BrowserRouter>
  );
}