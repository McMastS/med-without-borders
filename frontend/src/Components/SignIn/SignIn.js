import React from 'react';
import { Link } from 'react-router-dom';

class SignIn extends React.Component {
	constructor(props) {
		super(props);
		this.state = {
			signInEmail: '',
			signInPassword: '',
		}
	}

	onEmailChange = (event) => {
		this.setState({ signInEmail: event.target.value });
	}

	onPasswordChange = (event) => {
		this.setState({ signInPassword: event.target.value });
	}

	onSubmitSignIn = () => {
		fetch('http://api.medwithoutborders.org/source/login_normal', {
			method: 'post',
			headers: {'Content-Type': 'application/json'},
			body: JSON.stringify({
				username: this.state.signInEmail,
				password: this.state.signInPassword,
			})
		})
			.then(response => response.json())
			.then(user => {
				if (user.id) {
					localStorage.setItem('uuid', user.uuid);
					localStorage.setItem('token', user.token);
					this.props.history.push("/need");
				}
			})
	}

	render() {
		return (
				<article className="br3 ba dark-gray b--black-10 mv4 w-100 w-50-m w-25-l mw6 center shadow-5">
					<main className="pa4 black-80">
					  <div className="measure">
					    <fieldset id="sign_up" className="ba b--transparent ph0 mh0">
					      <legend className="f1 fw6 ph0 mh0">Sign In</legend>
					      <div className="mt3">
					        <label className="db fw6 lh-copy f6" htmlFor="email-address">Username</label>
					        <input className="pa2 input-reset ba bg-transparent w-100" type="email" name="email-address"  id="email-address" onChange={this.onEmailChange} />
					      </div>
					      <div className="mv3">
					        <label className="db fw6 lh-copy f6" htmlFor="password">Password</label>
					        <input className="b pa2 input-reset ba bg-transparent w-100" type="password" name="password"  id="password" onChange={this.onPasswordChange} />
					      </div>
					    </fieldset>
					    <div className="">
					      <input onClick={this.onSubmitSignIn} className="b ph3 pv2 input-reset ba b--black bg-transparent grow pointer f6 dib" type="submit" value="Sign in" />
					    </div>
					    <div className="lh-copy mt3">
					      <Link to="/register"><p className="f6 link dim black db pointer">Register</p></Link>
					      <Link to="/"><p className="f6 link dim black db pointer">Home</p></Link>
					    </div>
					  </div>
					</main>
				</article>
			)
		}
}

export default SignIn;