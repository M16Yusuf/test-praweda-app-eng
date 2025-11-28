import { BrowserRouter, Routes, Route } from "react-router";

import DataUserPage from "./page/DataUserPage";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route path="/" index element={<DataUserPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
