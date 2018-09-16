import React from 'react';
import { BrowserRouter, Link, Route, Switch } from 'react-router-dom';
import Homepage from '../Homepage/Homepage.js';
import Inventory from '../Inventory/Inventory.js';
import Supply from '../Supply/Supply.js';
import SignIn from '../SignIn/SignIn.js';
import Register from '../Register/Register.js';
import AboutPage from '../AboutPage/AboutPage.js'

const Main = () => {
	return (
		<Switch>
			<Route exact path='/' component={Homepage} />
			<Route path='/need' component={Inventory} />
			<Route path='/have' component={Supply} />
			<Route path='/signin' component={SignIn} />
			<Route path='/register' component={Register} />
			<Route path='/about' component={AboutPage} />
		</Switch>
	)
}

export default Main;