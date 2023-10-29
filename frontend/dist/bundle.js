(()=>{"use strict";var t={200:function(t,e,n){var r,o=this&&this.__extends||(r=function(t,e){return r=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(t,e){t.__proto__=e}||function(t,e){for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&(t[n]=e[n])},r(t,e)},function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Class extends value "+String(e)+" is not a constructor or null");function n(){this.constructor=t}r(t,e),t.prototype=null===e?Object.create(e):(n.prototype=e.prototype,new n)});Object.defineProperty(e,"__esModule",{value:!0}),e.EventList=void 0;var i=n(820),c=n(363),l=function(t){function e(){return null!==t&&t.apply(this,arguments)||this}return o(e,t),e.prototype.render=function(){var t=this.props.events.map((function(t){return c.createElement(i.EventListItem,{event:t})}));return c.createElement("table",{className:"table"},c.createElement("thead",null,c.createElement("tr",null,c.createElement("th",null,"Event"),c.createElement("th",null,"Where"),c.createElement("th",{colSpan:2},"When (start/end)"),c.createElement("th",null,"Actions"))),c.createElement("tbody",null,t))},e}(c.Component);e.EventList=l},68:function(t,e,n){var r,o=this&&this.__extends||(r=function(t,e){return r=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(t,e){t.__proto__=e}||function(t,e){for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&(t[n]=e[n])},r(t,e)},function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Class extends value "+String(e)+" is not a constructor or null");function n(){this.constructor=t}r(t,e),t.prototype=null===e?Object.create(e):(n.prototype=e.prototype,new n)});Object.defineProperty(e,"__esModule",{value:!0}),e.EventListContainer=void 0;var i=n(363),c=n(200),l=function(t){function e(e){var n=t.call(this,e)||this;return n.state={loading:!0,events:[]},fetch(e.eventListURL).then((function(t){return t.json()})).then((function(t){n.setState({loading:!1,events:t})})),n}return o(e,t),e.prototype.render=function(){return this.state.loading?i.createElement("div",null,"Loading..."):i.createElement(c.EventList,{events:this.state.events})},e}(i.Component);e.EventListContainer=l},820:function(t,e,n){var r,o=this&&this.__extends||(r=function(t,e){return r=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(t,e){t.__proto__=e}||function(t,e){for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&(t[n]=e[n])},r(t,e)},function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Class extends value "+String(e)+" is not a constructor or null");function n(){this.constructor=t}r(t,e),t.prototype=null===e?Object.create(e):(n.prototype=e.prototype,new n)});Object.defineProperty(e,"__esModule",{value:!0}),e.EventListItem=void 0;var i=n(363),c=function(t){function e(){return null!==t&&t.apply(this,arguments)||this}return o(e,t),e.prototype.render=function(){var t=new Date(1e3*this.props.event.StartDate),e=new Date(1e3*this.props.event.EndDate);return i.createElement("tr",null,i.createElement("td",null,this.props.event.Name),i.createElement("td",null,this.props.event.Location.Name),i.createElement("td",null,t.toLocaleDateString()),i.createElement("td",null,e.toLocaleTimeString()),i.createElement("td",null))},e}(i.Component);e.EventListItem=c},629:function(t,e,n){var r,o=this&&this.__extends||(r=function(t,e){return r=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(t,e){t.__proto__=e}||function(t,e){for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&(t[n]=e[n])},r(t,e)},function(t,e){if("function"!=typeof e&&null!==e)throw new TypeError("Class extends value "+String(e)+" is not a constructor or null");function n(){this.constructor=t}r(t,e),t.prototype=null===e?Object.create(e):(n.prototype=e.prototype,new n)});Object.defineProperty(e,"__esModule",{value:!0});var i=n(363),c=n(533),l=n(68),a=function(t){function e(){return null!==t&&t.apply(this,arguments)||this}return o(e,t),e.prototype.render=function(){return i.createElement("div",{className:"container"},i.createElement("h1",null,"MyEvents"),i.createElement(l.EventListContainer,{eventListURL:"http://localhost:8181"}))},e}(i.Component);c.render(i.createElement(a,null),document.getElementById("pEvents-app"))},363:t=>{t.exports=React},533:t=>{t.exports=ReactDOM}},e={};!function n(r){var o=e[r];if(void 0!==o)return o.exports;var i=e[r]={exports:{}};return t[r].call(i.exports,i,i.exports,n),i.exports}(629)})();