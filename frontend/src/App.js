import React, { useEffect } from 'react';
import './App.scss';
import {
  Switch,
  Route,
  Redirect,
  useLocation,
} from 'react-router-dom';
import Header from './components/Header'
import MainPage from './pages/main'
import PetitionPage from './pages/petition'

function App() {
  const location = useLocation()
  useEffect(() => {
    try {
      window.scrollTo(0, 0)
    } catch(e) {
      console.error(e)
    }
  }, [location.pathname])

  return (
    <>
      <Header />
      <div className='main-wrap'>
        <Switch>
          <Route path='/' exact component={MainPage} />
          <Route path='/petition' exact component={PetitionPage} />
          <Redirect to='/' />
        </Switch>
      </div>
    </>
  );
}

export default App;
