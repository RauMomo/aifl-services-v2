import EmberRouter from '@embroider/router';
import config from 'ember-website/config/environment';

export default class Router extends EmberRouter {
  location = config.locationType;
  rootURL = config.rootURL;
}

Router.map(function () {
  // Disguised Admin Authentication & Dashboard Routes
  this.route('portal-access', function () {
    this.route('dashboard');
  });
});
