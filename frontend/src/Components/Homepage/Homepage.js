import React from 'react';
import './Homepage.css';
import About from './About.js';
import { Link } from 'react-router-dom';

const Homepage = () => {
  const checkSignIn = () => {
  	const uuid = localStorage.getItem('uuid');
  	const token = localStorage.getItem('token');

  	if (uuid != null && token != null) {
  		return true;
  	}
  	return false;
  }

  return (
	  <div className="tc">
	    <header className="sans-serif">
				  <div className="cover bg-left bg-center-l background">
				    <div className="bg-black-80 pb5 pb6-m pb7-l">
				      <nav className="dt w-100 mw8 center"> 
				        <div className="dtc w2 v-mid pa3">
				          <a href="/" className="dib w2 h2 pa1 ba b--white-90 grow-large border-box">
				            <svg className="link white-90 hover-white skull" data-icon="skull" viewBox="0 0 32 32"><title>skull icon</title><path d="M16 0 C6 0 2 4 2 14 L2 22 L6 24 L6 30 L26 30 L26 24 L30 22 L30 14 C30 4 26 0 16 0 M9 12 A4.5 4.5 0 0 1 9 21 A4.5 4.5 0 0 1 9 12 M23 12 A4.5 4.5 0 0 1 23 21 A4.5 4.5 0 0 1 23 12"></path></svg>
				          </a>
				        </div>
				        <div className="dtc v-mid tr pa3">
				          <a className="f5 fw4 hover-white no-underline white-70 dn dib-ns pv2 ph3" href="/" >Donate</a> 
				          <Link to="/about"><a className="f5 fw4 hover-white no-underline white-70 dn dib-l pv2 ph3" href="/" >About</a> </Link>
				          <Link to="/signin"><a className="f5 fw4 hover-white no-underline white-70 dn dib-l pv2 ph3" href="/" >Sign In</a> </Link>
				          {
				          	checkSignIn() ? <Link to="/register"><a className="f5 fw4 hover-white no-underline white-70 dib ml2 pv2 ph3 ba" href="/" >Sign Up</a></Link> : <a className="f5 fw4 hover-white no-underline white-70 dib ml2 pv2 ph3 ba" href="/" >Sign Out</a>
				          }
				          
				        </div>
				      </nav> 
				      <div className="tc-l mt4 mt5-m mt6-l ph3">
				        <h1 className="f2 f1-l fw2 white-90 mb0 lh-title">Medicine Without Borders</h1>
				        <h2 className="fw1 f3 white-80 mt3 mb4">Connecting medicine supply with demand, in emergency situations</h2>
				        <Link to="/need"><a className="f6 no-underline grow dib v-mid white bg-red ba b--red ph3 pv2 mb3" href="">I need medicine</a></Link>
				        <span className="dib v-mid ph3 white-70 mb3">or</span>
				        <Link to="/have"><a className="f6 no-underline grow dib v-mid white ba b--white ph3 pv2 mb3" href="">I have medicine</a></Link>
				      </div>
				    </div>
				  </div> 
				</header>
				<About />
			</div>
  )
}

export default Homepage;
