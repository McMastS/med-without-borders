import React from 'react';
import Navbar from '../Navbar/Navbar.js';
import './Supply.css';

const Supply = () => {
	const onSubmit = () => {
		fetch('medwithoutborders.org/medicine')
	}

	return (
		<div>
			<Navbar />
			<div className="tc">
				<h1>Please input your inventory levels for the following medicines:</h1>
				<form className="br3 ba dark-gray b--black-10 mv4 mw6 center shadow-5">
					<label>Aspirin</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Aspirin'/>
					<label>Amiloride</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Amiloride'/>
					<label>Amiodarone</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Amiodarone'/>
					<label>Bisoprolol</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Bisoprolol'/>
					<label>Clopidogrel</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Clopidogrel'/> 
					<label>Digoxin</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Digoxin'/>
					<label>Furosemide</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Furosemide'/>
					<label>Losartan</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Losartan'/>
					<label>Methyldopa</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Methyldopa'/>
					<label>Nifedipine</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Nifedipine'/>
					<label>Spironolactone</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Spironolactone'/>
					<label>Streptokinase</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Streptokinase'/>
					<label>Verapamil</label>
					<input className="b pa2 input-reset ba bg-transparent w-30"type="text" name='Verapamil'/>
					<div className="">
					      <input onClick={this.onFormSubmit} className="b ph3 pv2 input-reset ba b--black bg-transparent pointer f6 dib red" type="submit" value="Submit" />
					</div>

				</form>
			</div>
		</div>
	);
}

export default Supply;