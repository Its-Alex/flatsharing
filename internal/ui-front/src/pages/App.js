import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import '../css/App.css'

import Navbar from '../components/Navbar'
import Login from './Login'

const App = () => {
  return (
    <Router>
      <Navbar />
      <div>
        <Route path='/' component={Login} />
      </div>
    </Router>
  )
}

export default App
