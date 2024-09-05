
import { Outlet } from 'react-router-dom'
import NavBar from './components/NavBar'
import Footer from './components/Footer'
import HeaderFooter from './components/HeaderFooter'

function App() {

  return (
    <>
            <NavBar />
            <Outlet/>
            <HeaderFooter/>
            <Footer/>
    </>
  )
}

export default App
