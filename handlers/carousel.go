package handlers

import "strconv"

// Discord shows at most 4 images per link preview, and a preview with two or
// more og:image tags is rendered as a gallery. A carousel is therefore split
// into at most 4 grids instead of one very tall strip.
const (
	maxPreviewImages = 4
	// A single grid never holds more than this many images, so a huge carousel
	// still produces readable grids (and the rest is simply not shown).
	maxImagesPerGrid = 9
)

// gridPages returns how many grids a carousel holding `total` images needs.
// It returns 0 when the carousel is small enough to be shown image by image.
func gridPages(total int) int {
	if total <= maxPreviewImages {
		return 0
	}

	pages := (total + maxPreviewImages - 1) / maxPreviewImages
	if pages > maxPreviewImages {
		pages = maxPreviewImages
	}

	// Spread the images over the pages, but never past maxImagesPerGrid.
	if (total+pages-1)/pages > maxImagesPerGrid {
		pages = (total + maxImagesPerGrid - 1) / maxImagesPerGrid
		if pages > maxPreviewImages {
			pages = maxPreviewImages
		}
	}
	return pages
}

// gridChunkSize returns how many images go into each grid of a carousel holding
// `total` images.
func gridChunkSize(total int) int {
	pages := gridPages(total)
	if pages == 0 {
		return total
	}

	chunk := (total + pages - 1) / pages
	if chunk > maxImagesPerGrid {
		chunk = maxImagesPerGrid
	}
	return chunk
}

// carouselImageURLs returns the preview images of a carousel: one URL per image
// when there are few of them, otherwise one grid URL per group.
func carouselImageURLs(postID string, total int) []string {
	pages := gridPages(total)
	if pages == 0 {
		urls := make([]string, 0, total)
		for i := 1; i <= total; i++ {
			urls = append(urls, "/images/"+postID+"/"+strconv.Itoa(i))
		}
		return urls
	}

	urls := make([]string, 0, pages)
	for page := 1; page <= pages; page++ {
		urls = append(urls, "/grid/"+postID+"/"+strconv.Itoa(page))
	}
	return urls
}
