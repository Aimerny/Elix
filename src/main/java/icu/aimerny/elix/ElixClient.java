package icu.aimerny.elix;

import icu.aimerny.elix.registry.ModBlock;
import net.fabricmc.api.ClientModInitializer;
import net.fabricmc.fabric.api.blockrenderlayer.v1.BlockRenderLayerMap;
import net.minecraft.client.render.RenderLayer;

public class ElixClient implements ClientModInitializer {
    @Override
    public void onInitializeClient() {
        BlockRenderLayerMap.INSTANCE.putBlock(ModBlock.OBSIDIAN_GLASS_BLOCK, RenderLayer.getTranslucent());
    }
}
