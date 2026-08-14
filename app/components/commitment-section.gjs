import Component from '@glimmer/component';
import UiBadge from './ui-badge';

export default class CommitmentSection extends Component {
  <template>
    <section class="section-wrapper">
      <div class="commitment-card">
        
        <UiBadge @variant="fill">OUR COMMITMENT</UiBadge>
        
        <h2 class="commitment-quote">
          "Solusi digital terbaik bukan selalu solusi yang paling kompleks. Solusi terbaik adalah solusi yang benar-benar membantu organisasi bekerja lebih baik."
        </h2>

        <p class="commitment-subtext">
          AIFL berkomitmen menjadi partner teknologi terpercaya yang membantu organisasi Anda:
        </p>

        <div class="commitment-pills">
          <div class="commitment-pill">
            Memahami Prosesnya
          </div>
          <div class="commitment-pill">
            Memperbaiki Cara Kerjanya
          </div>
          <div class="commitment-pill teal">
            Bertumbuh Dengan Teknologi Tepat
          </div>
        </div>
      </div>
    </section>
  </template>
}
