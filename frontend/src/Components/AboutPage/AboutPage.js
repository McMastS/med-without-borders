import React from 'react';
import Navbar from '../Navbar/Navbar.js';
import './AboutPage.css';

const AboutPage = () => {
	return (
		<div>
			<Navbar />
			<article className="cf ph3 ph5-ns pv5">
			  <header className="fn fl-ns w-50-ns pr4-ns">
			    <h1 className="f2 lh-title fw9 mb3 mt0 pt3 bt bw2 red-text">
			      About Medicine Without Borders
			    </h1>
			    <h2 className="f3 mid-gray lh-title red-text">
			      Our story
			    </h2>
			    <time className="f6 ttu tracked gray red-text">2018 - 9 - 16</time>
			  </header>
			  <div className="fn fl-ns w-50-ns">
				    <p className="f4 lh-copy measure mt0-ns">
				      Med Without Borders is a non-profit, self-sustaining web service that addresses 
				      the lack of essential medicines globally and particularly in developing countries. 
				      By facilitating swift access to necessary medications in a regulated and secure 
				      environment, our surface helps alleviate the health of patients and saves more lives 
				      each day.
				    </p>
				    <p className="f4 lh-copy measure">
				      To ensure maximum transparency and patient safety, only internationally accredited 
				      hospitals as Joint Commission International or JCI-accredited hospitals are allowed 
				      to make orders through this service. Additionally, the offered essential medicines are
				      consistently reviewed and updated with the WHO Model List of Essential Medicines updates.
				      Our service relies on fast courier services such as, FedEx International Next Flight Service,
				      to secure delivery in under 24 hours.
				    </p>
			  	</div>
				</article>
			</div>
	)
}

export default AboutPage;