# CIN114 - Cinema Production Company Website

A professional portfolio website for **CIN114 (Cinema-114)**, a film production company that creates films, music videos, and videography services. Built with Next.js 14 and deployed with Docker.

## 🎬 About CIN114

CIN-114 is a production company built by artists, for artists. We specialize in:

- **Cinema**: Films worth a trip to the theatre
- **Music Videos**: Collaborative visual storytelling for musical artists
- **Videography**: High-end event recording and commercial production

Visit us at [cin114.net](https://cin114.net)

## ✨ Features

- **Film Catalog**: Showcase of completed films with detailed credits and links
- **Video Portfolio**: Music videos and commercial videography work
- **Responsive Design**: Optimized for desktop and mobile viewing
- **Contact System**: Professional contact form for inquiries
- **Social Media Integration**: Links to YouTube, Instagram, X, Patreon, and Vimeo
- **SEO Optimized**: Proper metadata and Open Graph tags

## 🛠️ Tech Stack

- **Frontend**: Next.js 14, React 18, TypeScript
- **Styling**: SCSS/Sass with custom components
- **Deployment**: Docker with NGINX reverse proxy
- **SSL**: SSL certificate support for HTTPS
- **Build Tools**: ESLint, TypeScript compiler

## 🚀 Getting Started

### Prerequisites

- Node.js 18+ (for local development)
- Docker and Docker Compose (for deployment)
- SSL certificates (for production)

### Local Development

1. **Clone the repository**
   ```bash
   git clone https://github.com/Identityofsine/cin114.git
   cd cin114
   ```

2. **Install dependencies**
   ```bash
   npm install
   ```

3. **Run the development server**
   ```bash
   npm run dev
   ```

4. **Open your browser**
   ```
   http://localhost:3000
   ```

### Available Scripts

- `npm run dev` - Start development server
- `npm run build` - Build for production
- `npm run start` - Start production server
- `npm run lint` - Run ESLint

## 🐳 Docker Deployment

### Development Environment

```bash
# Start development environment
docker-compose -f docker-compose.dev.yaml up -d
```

### Production Environment

1. **SSL Setup**
   - Refer to the [SSL documentation](./ssl/README.md) for certificate setup
   - Ensure DNS is configured for your domain

2. **Build and deploy**
   ```bash
   docker-compose up -d
   ```

3. **Access the website**
   ```
   https://your-domain.com
   ```

### Docker Configuration

The application uses a multi-service setup:
- **NextJS Container**: Runs the Next.js application
- **NGINX Container**: Reverse proxy with SSL termination
- **Shared Network**: Enables container communication

## 📁 Project Structure

```
cin114/
├── src/
│   ├── app/                 # Next.js app router
│   │   ├── (root)/         # Main pages
│   │   ├── catalog/        # Film catalog pages
│   │   └── styles/         # Page-specific styles
│   ├── components/         # Reusable React components
│   ├── services/           # Business logic
│   ├── types/              # TypeScript type definitions
│   └── util/               # Utility functions
├── public/                 # Static assets
│   ├── film/              # Film-related images
│   ├── home/              # Homepage assets
│   └── ui/                # UI icons and graphics
├── nginx/                 # NGINX configuration
├── ssl/                   # SSL certificates
└── docker-compose.yaml    # Production deployment
```

## 🎨 Content Management

### Adding New Films

1. Update `src/film.settings.ts` with new film metadata
2. Add film assets to `public/film/[film-name]/`
3. Create a new page in `src/app/(root)/catalog/[film-name]/`

### Updating Brand Information

- Modify `src/brand.settings.ts` for contact info and social links
- Update `src/app/template_metadata.ts` for SEO metadata

## 🌐 Social Media

- **YouTube**: [@CIN114](https://www.youtube.com/@CIN114)
- **Instagram**: [@cin114films](https://www.instagram.com/cin114films)
- **X (Twitter)**: [@CIN114films](https://x.com/CIN114films)
- **Patreon**: [CIN114](https://www.patreon.com/CIN114)
- **Vimeo**: [CIN114](https://vimeo.com/user223422993)

## 📞 Contact

For business inquiries: [contact@cin114.net](mailto:contact@cin114.net)

## 🔧 Development Notes

- Uses Next.js 14 with App Router
- TypeScript for type safety
- SCSS modules for styling
- Docker for consistent deployment
- NGINX for production-grade serving

## 📝 License

This project is private and proprietary to CIN114.
