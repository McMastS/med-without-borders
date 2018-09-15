import React from 'react';
import './Register.css';

class Register extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			username: '',
			password: '',
			name: '',
			address: '',
			phone_number: '',
			photo_url: '',
			initial_inventory: '',
			prices: '',
		}
	}

	onNameChange = (event) => {
		this.setState({ name: event.target.value });
	}

	onUserChange = (event) => {
		this.setState({ username: event.target.value });
	}

	onPasswordChange = (event) => {
		this.setState({ password: event.target.value });
	}

	onAddressChange = (event) => {
		this.setState({ address: event.target.value });
	}

	onNumberChange = (event) => {
		this.setState({ phone_number: event.target.value });
	}

	onPhotoChange = (event) => {
		this.setState({ photo_url: event.target.value });
	}
	
	onInventoryChange = (event) => {
		this.setState({ inventory: event.target.value });
	}

	onPricesChange = (event) => {
		this.setState({ prices: event.target.value });
	}


	onSubmitSignIn = () => {
		fetch('http://api.medwithoutborders.org/source/new_user', {
			method: 'post',
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify({
				username: this.state.username,
				password: this.state.password,
				name: this.state.name,
				address: this.state.address,
				phone_number: this.state.phone_number,
				photo_url: this.state.photo_url,
				inventory: this.state.inventory,
				prices: this.state.prices,
			})
		})
			.then(response => response.json())
			.then(user => {
				if (user.id) {
					this.props.loadUser(user)
					this.props.onRouteChange('home')
				}
			})
		
	}
	
	render() {
		return (
			<article className="br3 ba dark-gray b--black-10 mv4 w-100 w-50-m w-25-l mw6 center shadow-5">
				<main className="pa4 black-80">
				  <div className="measure">
				    <fieldset id="sign_up" className="ba b--transparent ph0 mh0">
				      <legend className="f1 fw6 ph0 mh0">Register</legend>
				      <div className="">
				        <label className="db fw6 lh-copy f6" htmlFor="name">Name</label>
				        <input className="pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="name"  id="name" onChange={this.onNameChange} />
				      </div>
				      <div className="mt3">
				        <label className="db fw6 lh-copy f6" htmlFor="email-address">Username</label>
				        <input className="pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="email" name="email-address"  id="email-address" onChange={this.onUserChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="password">Password</label>
				        <input className="b pa2 input-reset ba bg-transparent w-100" type="password" name="password"  id="password" onChange={this.onPasswordChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="address">Address</label>
				        <input className="b pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="address"  id="address" onChange={this.onAddressChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="phone_number">Phone Number</label>
				        <input className="b pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="phone_number"  id="phone_number" onChange={this.onNumberChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="photo_url">Photo Url</label>
				        <input className="b pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="photo_url"  id="photo_url" onChange={this.onPhotoChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="inventory">Existing Inventory</label>
				        <p>Enter your inventory in standard units for the following drugs in order seperated by spaces: </p>
				        <p>Aspirin Amiloride Amiodarone Bisoprolol Clopidogrel Digoxin Furosemide Losartan Methyldopa Nifedipine Spironolactone Streptokinase Verapamil</p>
				        <input className="b pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="inventory"  id="inventory" onChange={this.onInventoryChange} />
				      </div>
				      <div className="mv3">
				        <label className="db fw6 lh-copy f6" htmlFor="prices">Prices</label>
				        <p>Enter your price in $ for the following drugs in order seperated by spaces: </p>
				        <p>Aspirin Amiloride Amiodarone Bisoprolol Clopidogrel Digoxin Furosemide Losartan Methyldopa Nifedipine Spironolactone Streptokinase Verapamil</p>
				        <input className="b pa2 input-reset ba bg-transparent hover-bg-black hover-white w-100" type="text" name="prices"  id="prices" onChange={this.onPricesChange} />
				      </div>
				    </fieldset>
				    <div className="">
				      <input onClick={this.onSubmitSignIn} className="b ph3 pv2 input-reset ba b--black bg-transparent grow pointer f6 dib" type="submit" value="Register" />
				    </div>
				  </div>
				</main>
			</article>
		)
	}		
}

export default Register;