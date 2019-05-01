import React from 'react'
import { BrowserRouter as Router, Route } from 'react-router-dom'
import '../css/App.css'

import Navbar from '../components/Navbar'
import Home from './Home'

function App () {
  return (
    <Router>
      <Navbar />
      <div>
        <Route path='/' component={Home} />
      </div>
    </Router>
  )
}

export default App
