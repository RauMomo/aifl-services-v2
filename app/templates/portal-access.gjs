import Component from '@glimmer/component';
import Navbar from '../components/navbar';
import FooterSection from '../components/footer-section';

export default class PortalAccessTemplate extends Component {
  <template>
    <Navbar />

    <div class="portal-route-content">
      {{outlet}}
    </div>

    <FooterSection />
  </template>
}
