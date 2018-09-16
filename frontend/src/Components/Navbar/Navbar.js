import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = () => {
	return (
		<nav className="flex justify-between bb bg-red b--white-10">
	  <a className="link white-70 hover-white no-underline flex items-center pa3" href="">
	    <svg
	      className="dib h1 w1"
	      data-icon="grid"
	      viewBox="0 0 32 32">
	      <title>I haven't slept in 2 days :)</title>
	    </svg>
	  </a>
	  <div className="flex-grow pa3 flex items-center">
	    <Link to="/"><a className="f6 link dib white dim mr3 mr4-ns" href="#0">Home</a></Link>
	    <a className="f6 dib white bg-animate hover-bg-white hover-black no-underline pv2 ph4 br-pill ba b--white-20" href="#0">Sign In</a>
	  </div>
	</nav>
	)
	 
}

export default Navbar;