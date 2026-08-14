import Component from '@glimmer/component';
import Navbar from '../components/navbar';
import HeroSection from '../components/hero-section';
import AboutSection from '../components/about-section';
import BeliefSection from '../components/belief-section';
import ProblemsSection from '../components/problems-section';
import ServicesSection from '../components/services-section';
import ProcessSection from '../components/process-section';
import CaseStudiesSection from '../components/case-studies-section';
import WhyAiflSection from '../components/why-aifl-section';
import CommitmentSection from '../components/commitment-section';
import ContactSection from '../components/contact-section';
import FooterSection from '../components/footer-section';

export default class IndexTemplate extends Component {
  <template>
    <Navbar />

    <main class="main-content">
      <HeroSection />
      <AboutSection />
      <BeliefSection />
      <ProblemsSection />
      <ServicesSection />
      <ProcessSection />
      <CaseStudiesSection />
      <WhyAiflSection />
      <CommitmentSection />
      <ContactSection />
    </main>

    <FooterSection />
  </template>
}
