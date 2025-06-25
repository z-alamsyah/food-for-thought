import { BrowserRouter, Route, Routes } from "react-router-dom"
import UserRegister from "./pages/UserRegister"
import Layout from "./components/Layout"

function App() {

  return (
    <BrowserRouter>
      <Routes >
        <Route element={<Layout />}>
          <Route path="/register" element={<UserRegister />} />
        </Route>
      </Routes>
    </BrowserRouter>     
  )
}

export default App
